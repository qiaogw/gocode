package gormx

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/mathx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

// 参见 doc/sql-cache.md
const cacheSafeGapBetweenIndexAndPrimary = time.Second * 5

// spanName 用于标识 SQL 执行的 span 名称。
const spanName = "sql"

// 使过期时间不稳定，以避免大量缓存项同时过期
// 使不稳定的过期时间在 [0.95, 1.05] * 秒之间
const expiryDeviation = 0.05

var (
	// ErrNotFound 是 gorm.ErrRecordNotFound 的别名。
	ErrNotFound = gorm.ErrRecordNotFound

	// 不能每个连接使用一个 SingleFlight，因为多个连接可能共享相同的缓存键。
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("gorm")
)

type (

	// ExecCtxFn 定义了 SQL 执行方法。
	ExecCtxFn func(conn *gorm.DB) error
	// IndexQueryCtxFn 定义了基于唯一索引的查询方法。
	IndexQueryCtxFn func(conn *gorm.DB, v interface{}) (interface{}, error)
	// PrimaryQueryCtxFn 定义了基于主键的查询方法。
	PrimaryQueryCtxFn func(conn *gorm.DB, v, primary interface{}) error
	// QueryCtxFn 定义了查询方法。
	QueryCtxFn func(conn *gorm.DB) error

	CachedConn struct {
		db                 *gorm.DB
		cache              cache.Cache
		unstableExpiryTime mathx.Unstable
	}

	Conn struct {
		db *gorm.DB
	}
)

// NewConn 返回一个带有 Redis 集群缓存的 CachedConn。
func NewConn(db *gorm.DB, c cache.CacheConf, opts ...cache.Option) CachedConn {
	cc := cache.New(c, singleFlights, stats, ErrNotFound, opts...)
	return NewConnWithCache(db, cc)
}

// NewConnWithCache 返回一个带有自定义缓存的 CachedConn。
func NewConnWithCache(db *gorm.DB, c cache.Cache) CachedConn {
	return CachedConn{
		db:                 db,
		cache:              c,
		unstableExpiryTime: mathx.NewUnstable(expiryDeviation),
	}
}

// NewNodeConn 返回一个带有 Redis 单节点缓存的 CachedConn。
func NewNodeConn(db *gorm.DB, rds *redis.Redis, opts ...cache.Option) CachedConn {
	cc := cache.NewNode(rds, singleFlights, stats, ErrNotFound, opts...)
	return NewConnWithCache(db, cc)
}

// DelCache 删除指定键的缓存。
func (cc CachedConn) DelCache(keys ...string) error {
	return cc.cache.DelCtx(context.Background(), keys...)
}

// DelCacheCtx 删除指定键的缓存。
func (cc CachedConn) DelCacheCtx(ctx context.Context, keys ...string) error {
	return cc.cache.DelCtx(ctx, keys...)
}

// GetCache 将给定键的缓存反序列化到 v 中。
func (cc CachedConn) GetCache(key string, v interface{}) error {
	return cc.cache.GetCtx(context.Background(), key, v)
}

// GetCacheCtx 将给定键的缓存反序列化到 v 中。
func (cc CachedConn) GetCacheCtx(ctx context.Context, key string, v interface{}) error {
	return cc.cache.GetCtx(ctx, key, v)
}

// Exec 在给定键上运行给定的 exec，并返回执行结果。
func (cc CachedConn) Exec(exec ExecCtxFn, keys ...string) error {
	return cc.ExecCtx(context.Background(), exec, keys...)
}

// ExecCtx 在给定键上运行给定的 exec，并返回执行结果。
func (cc CachedConn) ExecCtx(ctx context.Context, execCtx ExecCtxFn, keys ...string) error {
	err := execCtx(cc.db.WithContext(ctx))
	if err != nil {
		return err
	}
	if err := cc.DelCacheCtx(ctx, keys...); err != nil {
		return err
	}
	return nil
}

// ExecNoCache 运行给定的 SQL 语句，不影响缓存。
func (cc Conn) ExecNoCache(exec ExecCtxFn) error {
	return cc.ExecNoCacheCtx(context.Background(), exec)
}

// ExecNoCacheCtx 运行给定的 SQL 语句，不影响缓存。
func (cc Conn) ExecNoCacheCtx(ctx context.Context, execCtx ExecCtxFn) (err error) {
	ctx, span := startSpan(ctx, "ExecNoCache")
	defer func() {
		endSpan(span, err)
	}()
	return execCtx(cc.db.WithContext(ctx))
}

// ExecNoCache 运行给定的 SQL 语句，不影响缓存。
func (cc CachedConn) ExecNoCache(exec ExecCtxFn) error {
	return cc.ExecNoCacheCtx(context.Background(), exec)
}

// ExecNoCacheCtx 运行给定的 SQL 语句，不影响缓存。
func (cc CachedConn) ExecNoCacheCtx(ctx context.Context, execCtx ExecCtxFn) (err error) {
	ctx, span := startSpan(ctx, "ExecNoCache")
	defer func() {
		endSpan(span, err)
	}()
	return execCtx(cc.db.WithContext(ctx))
}

// QueryRowIndex 将给定键的缓存反序列化到 v 中。
func (cc CachedConn) QueryRowIndex(v interface{}, key string, keyer func(primary interface{}) string,
	indexQuery IndexQueryCtxFn, primaryQuery PrimaryQueryCtxFn) error {
	return cc.QueryRowIndexCtx(context.Background(), v, key, keyer, indexQuery, primaryQuery)
}

// QueryRowIndexCtx 将给定键的缓存反序列化到 v 中。
func (cc CachedConn) QueryRowIndexCtx(ctx context.Context, v interface{}, key string, keyer func(primary interface{}) string, indexQuery IndexQueryCtxFn, primaryQuery PrimaryQueryCtxFn) (err error) {
	ctx, span := startSpan(ctx, "QueryRowIndex")
	defer func() {
		endSpan(span, err)
	}()

	var primaryKey interface{}
	var found bool

	if err = cc.cache.TakeWithExpireCtx(ctx, &primaryKey, key, func(val interface{}, expire time.Duration) error {
		primaryKey, err = indexQuery(cc.db.WithContext(ctx), v)
		if err != nil {
			return err
		}
		found = true
		return cc.cache.SetWithExpireCtx(ctx, keyer(primaryKey), v, expire+cacheSafeGapBetweenIndexAndPrimary)
	}); err != nil {
		return err
	}
	if found {
		return nil
	}
	return cc.cache.TakeCtx(ctx, v, keyer(primaryKey), func(v interface{}) error {
		return primaryQuery(cc.db.WithContext(ctx), v, primaryKey)
	})
}

// QueryCtx 将给定键的缓存反序列化到 v 中。
func (cc CachedConn) QueryCtx(ctx context.Context, v interface{}, key string, query QueryCtxFn) (err error) {
	ctx, span := startSpan(ctx, "Query")
	defer func() {
		endSpan(span, err)
	}()
	return cc.cache.TakeCtx(ctx, v, key, func(v interface{}) error {
		return query(cc.db.WithContext(ctx))
	})
}

// QueryNoCacheCtx 运行给定的查询，不影响缓存。
func (cc CachedConn) QueryNoCacheCtx(ctx context.Context, query QueryCtxFn) (err error) {
	ctx, span := startSpan(ctx, "QueryNoCache")
	defer func() {
		endSpan(span, err)
	}()
	return query(cc.db.WithContext(ctx))
}

// QueryWithExpireCtx 将给定键的缓存反序列化到 v 中，并设置过期时间和查询函数。
func (cc CachedConn) QueryWithExpireCtx(ctx context.Context, v interface{}, key string, expire time.Duration, query QueryCtxFn) (err error) {
	ctx, span := startSpan(ctx, "QueryWithExpire")
	defer func() {
		endSpan(span, err)
	}()
	err = cc.cache.TakeCtx(ctx, v, key, func(v interface{}) error {
		return query(cc.db.WithContext(ctx))
	})
	if err != nil {
		return err
	}
	return cc.cache.SetWithExpireCtx(ctx, key, v, cc.aroundDuration(expire))
}

// QueryWithCallbackExpireCtx 将给定键的缓存反序列化到 v 中，并通过回调设置过期时间和查询函数。
func (cc CachedConn) QueryWithCallbackExpireCtx(ctx context.Context, v interface{}, key string, query QueryCtxFn, callback func(interface{}) time.Duration) (err error) {
	ctx, span := startSpan(ctx, "QueryWithCallbackExpire")
	defer func() {
		endSpan(span, err)
	}()
	err = cc.cache.TakeCtx(ctx, v, key, func(v interface{}) error {
		return query(cc.db.WithContext(ctx))
	})
	if err != nil {
		return err
	}
	if callback == nil {
		return cc.QueryCtx(ctx, v, key, query)
	}
	return cc.cache.SetWithExpireCtx(ctx, key, v, callback(v))
}

// aroundDuration 返回一个不稳定的持续时间。
func (cc CachedConn) aroundDuration(duration time.Duration) time.Duration {
	return cc.unstableExpiryTime.AroundDuration(duration)
}

// SetCache 将 v 设置到缓存中，使用给定的键。
func (cc CachedConn) SetCache(key string, v interface{}) error {
	return cc.cache.SetCtx(context.Background(), key, v)
}

// SetCacheCtx 将 v 设置到缓存中，使用给定的键。
func (cc CachedConn) SetCacheCtx(ctx context.Context, key string, val interface{}) error {
	return cc.cache.SetCtx(ctx, key, val)
}

// SetCacheWithExpireCtx 将 v 设置到缓存中，使用给定的键和过期时间。
func (cc CachedConn) SetCacheWithExpireCtx(ctx context.Context, key string, val interface{}, expire time.Duration) error {
	return cc.cache.SetWithExpireCtx(ctx, key, val, expire)
}

// Transact 在事务模式下运行给定的 fn。
func (cc CachedConn) Transact(fn func(db *gorm.DB) error, opts ...*sql.TxOptions) error {
	return cc.TransactCtx(context.Background(), fn, opts...)
}

// TransactCtx 在事务模式下运行给定的 fn。
func (cc CachedConn) TransactCtx(ctx context.Context, fn func(db *gorm.DB) error, opts ...*sql.TxOptions) error {
	return cc.db.WithContext(ctx).Transaction(fn, opts...)
}

var sqlAttributeKey = attribute.Key("sql.method")

// startSpan 启动一个新的 span。
func startSpan(ctx context.Context, method string) (context.Context, oteltrace.Span) {
	tracer := otel.Tracer(trace.TraceName)
	start, span := tracer.Start(ctx,
		spanName,
		oteltrace.WithSpanKind(oteltrace.SpanKindClient),
	)
	span.SetAttributes(sqlAttributeKey.String(method))

	return start, span
}

// endSpan 结束 span 并记录错误。
func endSpan(span oteltrace.Span, err error) {
	defer span.End()

	if err == nil || errors.Is(err, ErrNotFound) {
		span.SetStatus(codes.Ok, "")
		return
	}

	span.SetStatus(codes.Error, err.Error())
	span.RecordError(err)
}

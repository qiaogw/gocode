package batchx

import (
	"context"
	"github.com/qiaogw/gocode/common/gormx"
	"gorm.io/gorm"
)

// BatchExecModel 定义了批量执行操作的接口，要求实现获取缓存键和在上下文中执行数据库操作
type BatchExecModel[DBModel any] interface {
	// GetCacheKeys 根据传入的数据返回对应的缓存键列表
	GetCacheKeys(data *DBModel) []string
	// ExecCtx 在给定的上下文中执行数据库操作，并传入可选的缓存键列表
	ExecCtx(ctx context.Context, execCtx gormx.ExecCtxFn, keys ...string) error
}

// BatchExecCtx 对批量数据执行操作，并自动处理缓存键
// 参数说明：
//
//	ctx：上下文
//	model：实现 BatchExecModel 接口的模型
//	olds：旧数据列表，基于该列表生成缓存键
//	exec：具体的数据库操作函数
//
// 如果 olds 为空，则直接返回 nil；否则先获取缓存键，再通过 ExecCtx 执行操作
func BatchExecCtx[DBModel any, Model BatchExecModel[DBModel]](ctx context.Context, model Model, olds []DBModel, exec func(db *gorm.DB) error) error {
	if len(olds) == 0 {
		return nil
	}
	// 根据旧数据获取所有相关的缓存键
	cacheKeys := getCacheKeysByMultiData(model, olds)

	// 在给定的上下文中执行数据库操作，并传入缓存键
	err := model.ExecCtx(ctx, func(conn *gorm.DB) error {
		return exec(conn)
	}, cacheKeys...)
	return err
}

// getCacheKeysByMultiData 根据数据列表生成所有缓存键，并进行去重
// 参数说明：
//
//	m：实现 BatchExecModel 接口的模型
//	data：数据列表
//
// 如果数据列表为空，则返回空的字符串切片；否则对每个数据调用 GetCacheKeys 并合并所有键，再去重后返回
func getCacheKeysByMultiData[DBModel any, Model BatchExecModel[DBModel]](m Model, data []DBModel) []string {
	if len(data) == 0 {
		return []string{}
	}
	var keys []string
	// 遍历每个数据，收集缓存键
	for _, v := range data {
		keys = append(keys, m.GetCacheKeys(&v)...)
	}
	// 去重后返回缓存键列表
	keys = uniqueKeys(keys)
	return keys
}

// uniqueKeys 对给定的字符串切片进行去重，返回唯一的键列表
func uniqueKeys(keys []string) []string {
	keySet := make(map[string]struct{})
	// 将每个键存入 map 以实现去重
	for _, key := range keys {
		keySet[key] = struct{}{}
	}

	uniKeys := make([]string, 0, len(keySet))
	// 从 map 中提取唯一键
	for key := range keySet {
		uniKeys = append(uniKeys, key)
	}

	return uniKeys
}

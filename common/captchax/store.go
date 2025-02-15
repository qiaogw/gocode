package captchax

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"time"
)

// cacheStore 实现了 base64Captcha.Store 接口，基于 go-zero 的缓存实现验证码存储
type cacheStore struct {
	cache      cache.Cache   // 缓存对象
	expiration time.Duration // 过期时长
}

// NewCacheStore 返回一个新的验证码内存存储器，使用指定的缓存对象和过期时间。
// 返回的存储器需要通过 SetCustomStore 注册，以替换默认的存储器。
func NewCacheStore(cache cache.Cache, expiration time.Duration) base64Captcha.Store {
	s := new(cacheStore)
	s.cache = cache
	s.expiration = expiration
	return s
}

// Set 将验证码值存储到缓存中，并设置对应的过期时间。
func (e *cacheStore) Set(id string, value string) error {
	err := e.cache.SetWithExpire(id, value, e.expiration)
	return err
}

// Get 根据验证码 id 获取存储的验证码值。
// 参数 clear 指定是否在获取后删除该验证码记录。
// 如果成功获取则返回验证码，否则返回空字符串。
func (e *cacheStore) Get(id string, clear bool) string {
	var v string
	err := e.cache.Get(id, &v)
	if err == nil {
		if clear {
			_ = e.cache.Del(id)
		}
		return v
	}
	return ""
}

// Verify 直接验证验证码答案是否正确，内部调用 Get 方法进行比对。
func (e *cacheStore) Verify(id, answer string, clear bool) bool {
	return e.Get(id, clear) == answer
}

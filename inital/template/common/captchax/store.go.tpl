package captchax

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"time"
)

type cacheStore struct {
	cache      cache.Cache
	expiration time.Duration
}

// NewCacheStore returns a new standard memory store for captchas with the
// given collection threshold and expiration time (duration). The returned
// store must be registered with SetCustomStore to replace the default one.
func NewCacheStore(cache cache.Cache, expiration time.Duration) base64Captcha.Store {
	s := new(cacheStore)
	s.cache = cache
	s.expiration = expiration
	return s
}

// Set sets the digits for the captcha id.
func (e *cacheStore) Set(id string, value string) error {
	err := e.cache.SetWithExpire(id, value, e.expiration)
	return err
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
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

//Verify captcha answer directly
func (e *cacheStore) Verify(id, answer string, clear bool) bool {
	return e.Get(id, clear) == answer
}

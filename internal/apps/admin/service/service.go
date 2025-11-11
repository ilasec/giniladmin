package service

import (
	"context"
	"giniladmin/internal/apps/admin/config"
	"giniladmin/internal/apps/admin/service/auth"
	"giniladmin/internal/apps/admin/service/health"
	"giniladmin/internal/apps/admin/service/oauth"
	"giniladmin/internal/apps/admin/service/system"
	"giniladmin/pkg/utils"
	"giniladmin/pkg/utils/cache"
	"giniladmin/pkg/utils/jtoken"
	"github.com/allegro/bigcache/v3"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

var preInitList = []func(ctx context.Context){
	health.Init,
	auth.Init,
	system.Init,
	oauth.Init,
}

var serviceOnce sync.Once

//func initJwt(expires, key string) (jwt *middleware.GinJWTMiddleware, err error) {
//	dr, err := utils.ParseDuration(expires)
//	if err != nil {
//		return
//	}
//	jwt, err = middleware.NewDefaultJwt(key, dr.Milliseconds(), nil)
//	return
//}

func initJwt(c config.Jwt) (err error) {
	jtoken.SigningKey = []byte(c.SigningKey)
	jtoken.BufferTime = c.BufferTime
	jtoken.ExpiresTime = c.ExpiresTime
	jtoken.Issuer = c.Issuer
	jtoken.Concurrency = &singleflight.Group{}

	return
}

func initCache(evicted int) (cache *bigcache.BigCache) {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(time.Duration(evicted)*time.Second))
	utils.CheckAndExit(err)
	return
}

func initCache2(evicted string) (ret *cache.Cache) {
	d, err := utils.ParseDuration(evicted)
	utils.CheckAndExit(err)
	c := cache.NewCache(
		cache.SetDefaultExpire(d),
	)
	utils.CheckAndExit(err)
	ret = &c
	return
}

func Setup(ctx context.Context) error {
	serviceOnce.Do(func() {
		v := ctx.Value("value").(map[string]any)
		c := v["conf"].(*config.Config)

		//jtoken init
		err := initJwt(c.Jwt)
		utils.CheckAndExit(err)

		//blackcache init
		v["bcache"] = initCache2(c.Jwt.ExpiresTime)

		ctx := context.WithValue(context.Background(), "value", v)

		for _, f := range preInitList {
			f(ctx)
		}
	})

	return nil
}

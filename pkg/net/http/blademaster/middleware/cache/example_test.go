package cache_test

import (
	"time"

	"github.com/welcome112s/go-library/pkg/cache/memcache"
	"github.com/welcome112s/go-library/pkg/container/pool"
	"github.com/welcome112s/go-library/pkg/ecode"
	"github.com/welcome112s/go-library/pkg/net/http/blademaster"
	"github.com/welcome112s/go-library/pkg/net/http/blademaster/middleware/cache"
	"github.com/welcome112s/go-library/pkg/net/http/blademaster/middleware/cache/store"
	xtime "github.com/welcome112s/go-library/pkg/time"

	"github.com/pkg/errors"
)

// This example create a cache middleware instance and two cache policy,
// then attach them to the specified path.
//
// The `PageCache` policy will attempt to cache the whole response by URI.
// It usually used to cache the common response.
//
// The `Degrader` policy usually used to prevent the API totaly unavailable if any disaster is happen.
// A succeeded response will be cached per 600s.
// The cache key is generated by specified args and its values.
// You can using file or memcache as cache backend for degradation currently.
//
// The `Cache` policy is used to work with multilevel HTTP caching architecture.
// This will cause client side response caching.
// We only support weak validator with `ETag` header currently.
func Example() {
	mc := store.NewMemcache(&memcache.Config{
		Config: &pool.Config{
			Active:      10,
			Idle:        2,
			IdleTimeout: xtime.Duration(time.Second),
		},
		Name:         "test",
		Proto:        "tcp",
		Addr:         "172.16.33.54:11211",
		DialTimeout:  xtime.Duration(time.Second),
		ReadTimeout:  xtime.Duration(time.Second),
		WriteTimeout: xtime.Duration(time.Second),
	})
	ca := cache.New(mc)
	deg := cache.NewDegrader(10)
	pc := cache.NewPage(10)
	ctl := cache.NewControl(10)
	filter := func(ctx *blademaster.Context) bool {
		if ctx.Request.Form.Get("cache") == "false" {
			return false
		}
		return true
	}

	engine := blademaster.Default()
	engine.GET("/users/profile", ca.Cache(deg.Args("name", "age"), nil), func(c *blademaster.Context) {
		values := c.Request.URL.Query()
		name := values.Get("name")
		age := values.Get("age")

		err := errors.New("error from others") // error from other call
		if err != nil {
			// mark this response should be degraded
			c.JSON(nil, ecode.Degrade)
			return
		}
		c.JSON(map[string]string{"name": name, "age": age}, nil)
	})
	engine.GET("/users/index", ca.Cache(pc, nil), func(c *blademaster.Context) {
		c.String(200, "%s", "Title: User")
	})
	engine.GET("/users/list", ca.Cache(ctl, filter), func(c *blademaster.Context) {
		c.JSON([]string{"user1", "user2", "user3"}, nil)
	})
	engine.Run(":18080")
}

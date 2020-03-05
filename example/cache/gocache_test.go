package cache_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/tribunadigital/dataloaden/example"
	"github.com/tribunadigital/dataloaden/example/cache"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoCache(t *testing.T) {

	cacheConf := cache.UserLoaderGoCacheConfig{
		DefaultExpiration: 100 * time.Millisecond,
		CleanupInterval:   time.Minute,
	}
	c := 0
	dl := cache.NewUserLoader(cache.UserLoaderConfig{
		Wait: 10 * time.Microsecond,
		Fetch: func(keys []string) ([]*example.User, []error) {
			users := make([]*example.User, len(keys))
			errors := make([]error, len(keys))

			c++
			for i, key := range keys {
				users[i] = &example.User{
					ID:   key,
					Name: strconv.Itoa(c),
				}

			}

			return users, errors
		},
		Cache: cache.NewUserLoaderGoCache(cacheConf),
	})

	t.Run("Cache", func(t *testing.T) {
		// Load `U` and cache it
		u, err := dl.Load("U")
		require.NoError(t, err)

		atcache := u.Name

		// each Load incr `c`
		dl.Load("U")
		dl.Load("U")
		dl.Load("U")

		u, err = dl.Load("U")
		require.NoError(t, err)
		assert.Equal(t, atcache, u.Name)

		atcache = u.Name
		time.Sleep(cacheConf.DefaultExpiration)

		u, err = dl.Load("U")
		require.NoError(t, err)
		assert.NotEqual(t, atcache, u.Name)
		assert.Equal(t, u.Name, strconv.Itoa(c))
	})
}

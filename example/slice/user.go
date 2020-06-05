//go:generate ../../dataloaden UserSliceLoader string []github.com/tribunadigital/dataloaden/example.User

package slice

import (
	"context"
	"time"

	"github.com/tribunadigital/dataloaden/example"
)

func NewLoader() *UserSliceLoader {
	return &UserSliceLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(contexts []context.Context, keys []string) ([][]example.User, []error) {
			users := make([][]example.User, len(keys))
			errors := make([]error, len(keys))

			for i, key := range keys {
				users[i] = []example.User{{ID: key, Name: "user " + key}}
			}
			return users, errors
		},
	}
}

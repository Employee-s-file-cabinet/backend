package smap

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Employee-s-file-cabinet/backend/pkg/repoerr"
)

type smap struct {
	items sync.Map
	close chan struct{}
}

type item struct {
	data    int
	expires int64
}

// New creates sync map storage type of key-value.
// Simple, but thread-safe. Designed mainly for testing purposes.
func New(cleaningInterval time.Duration) *smap {
	smap := &smap{
		close: make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(cleaningInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				now := time.Now().UnixNano()

				smap.items.Range(func(key, value any) bool {
					item := value.(item)

					if item.expires > 0 && now > item.expires {
						smap.items.Delete(key)
					}

					return true
				})

			case <-smap.close:
				return
			}
		}
	}()

	return smap
}

// Set sets value by key.
func (smap *smap) Set(_ context.Context, key string, value int, duration time.Duration) error {
	var expires int64

	if duration > 0 {
		expires = time.Now().Add(duration).UnixNano()
	}

	smap.items.Store(key, item{
		data:    value,
		expires: expires,
	})

	return nil
}

// Get returns value by key.
func (smap *smap) Get(_ context.Context, key string) (int, error) {
	const op = "sync map: get"

	obj, exists := smap.items.Load(key)

	if !exists {
		return 0, repoerr.ErrRecordNotFound
	}

	item := obj.(item)

	if item.expires > 0 && time.Now().UnixNano() > item.expires {
		smap.items.Delete(key)
		return 0, fmt.Errorf("record not found")
	}

	return item.data, nil
}

// Delete deletes value by key.
func (smap *smap) Delete(_ context.Context, key string) error {
	smap.items.Delete(key)
	return nil
}

// Close closes storage: cleans internal map and stop cleaning work.
func (smap *smap) Close() error {
	smap.close <- struct{}{}
	smap.items = sync.Map{}

	return nil
}
package pokecache

import (
	"sync"
	"time"
)

type Cashe struct {
	cashe map[string]casheEntry
	mux   *sync.Mutex
}

type casheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCashe(interval time.Duration) Cashe {
	c := Cashe{
		cashe: make(map[string]casheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cashe) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cashe[key] = casheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cashe) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cashE, ok := c.cashe[key]
	return cashE.val, ok
}

func (c *Cashe) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cashe) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().Add(-interval)
	for k, v := range c.cashe {
		if v.createdAt.Before(timeAgo) {
			delete(c.cashe, k)
		}
	}
}

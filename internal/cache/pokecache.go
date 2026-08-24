package pokecache

import (
	"time"
	"sync"
	"github.com/dragonbitestail/pokedexcli/internal/logging"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	// duration time.Duration
	cacheMap map[string]cacheEntry
	m *sync.Mutex
}

var logr = ilogger.GetLogger()


func NewCache(duration time.Duration) *Cache {

	c := &Cache {
		cacheMap: make(map[string]cacheEntry),
		m: &sync.Mutex{},
	}

	go c.reapLoop(duration)
	logr.Info("pokecache >> NewCache() cache created", "duration", duration)
	return c

}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	defer c.m.Unlock()

	cEnt := cacheEntry {
		createdAt: time.Now(),
		val: val,
	}

	c.cacheMap[key] = cEnt

	logr.Info("pokecache >> Add() cache item added", "key", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	defer c.m.Unlock()

  cEntry, ok := c.cacheMap[key]
  if ! ok {
    return nil, false
  }
	logr.Info("pokecache >> Get() returning cache item", "key", key)

  return cEntry.val, true
}


func (c *Cache) reapLoop(d time.Duration) {
	tickr := time.NewTicker(d)

	for {
		<- tickr.C  // Block waiting for Ticker Channel to send signal every passed duration; d
		logr.Info("pokecache >> reapLoop() invalidating candidate cache items", "interval", d)
		c.m.Lock()
		for k, v := range c.cacheMap {
			if time.Now().Sub(v.createdAt) > d {
				delete(c.cacheMap, k)
			}
		}
		c.m.Unlock()
	}

}

package pokecache

import (
	"time"
	"sync"
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


func NewCache(duration time.Duration) *Cache {
	
	c := &Cache {
		cacheMap: make(map[string]cacheEntry),
		m: &sync.Mutex{},
	}

	go c.reapLoop(duration)
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

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	defer c.m.Unlock()

  if _, ok := c.cacheMap[key]; ! ok {
    return nil, false
  }
  
  return c.cacheMap[key].val, true
}


func (c *Cache) reapLoop(d time.Duration) {
	tickr := time.NewTicker(d)

	for {
		//t := <- tickr.C
		<- tickr.C
		c.m.Lock()
		for k, v := range c.cacheMap {
			if time.Now().Sub(v.createdAt) > d {
				delete(c.cacheMap, k)
			}
		}
		c.m.Unlock()
	}

}

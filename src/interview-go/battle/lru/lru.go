package lru

import (
	"container/list"
	"sync"
	"time"
)

type Cache struct {
	MaxEntries int
	OnEvicted  func(key Key, value interface{})
	ll         *list.List
	cache      map[interface{}]*list.Element
	ttl        time.Duration
	lock       sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators
type Key interface{}

type entry struct {
	key   Key
	value interface{}
	//if tll is nil, entry is not expire auto
	ttl *time.Time
}

func (e *entry) IsExpired() bool {
	if e.ttl == nil {
		return false
	}
	return time.Now().After(*e.ttl)
}

// New creates a new Cache.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func New(maxEntries int, ttl time.Duration, onEvicted func(key Key, value interface{})) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
		OnEvicted:  onEvicted,
		ttl:        ttl,
	}
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		//expired
		if ele.Value.(*entry).IsExpired() {
			c.removeElement(ele)
			return nil, false
		}
		ex := time.Now().Add(c.ttl)
		ele.Value.(*entry).ttl = &ex
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

// Put adds a value to the cache.
func (c *Cache) Put(key Key, value interface{}, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	var ex *time.Time = nil
	if ttl > 0 {
		expire := time.Now().Add(ttl)
		ex = &expire
	} else if c.ttl > 0 {
		expire := time.Now().Add(c.ttl)
		ex = &expire
	}
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value, ex})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.removeOldest()
	}
}

// Remove removes the provided key from the cache.
func (c *Cache) Remove(key Key) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.cache == nil {
		return false
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
		return true
	}
	return false
}

// removeOldest removes the oldest item from the cache.
func (c *Cache) removeOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

// Len returns the number of items in the cache.
func (c *Cache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

// Keys return all the keys in cache, from oldest to newest
func (c *Cache) Keys() []interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	keys := make([]interface{}, len(c.cache))
	i := 0
	for ent := c.ll.Back(); ent != nil; ent = ent.Prev() {
		keys[i] = ent.Value.(*entry).key
		i++
	}
	return keys
}

// Contains Check if a key exsists in cache without updating the recent-ness.
func (c *Cache) Contains(key interface{}) (ok bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if ent, ok := c.cache[key]; ok {
		if ent.Value.(*entry).IsExpired() {
			return false
		}
		return ok
	}
	return false
}

// Clear purges all stored items from the cache.
func (c *Cache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
}

package main

import "container/list"

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU(operators [][]int, k int) []int {
	// write code here
	var res []int
	cache := New(k)
	for _, opt := range operators {
		if opt[0] == 1 {
			cache.Set(opt[1], opt[2])
		} else {
			res = append(res, cache.Get(opt[1]))
		}
	}
	return res
}

type LRUCache struct {
	Capacity int
	list     *list.List
	cacheMap map[int]*list.Element
}
type CacheEntry struct {
	key   int
	value int
}

func New(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		list:     list.New(),
		cacheMap: make(map[int]*list.Element, cap),
	}
}
func (c *LRUCache) Set(key int, value int) {
	if ele, ok := c.cacheMap[key]; ok {
		c.list.MoveToFront(ele)
		ele.Value.(*CacheEntry).value = value
		return
	}
	newEle := c.list.PushFront(&CacheEntry{key: key, value: value})
	c.cacheMap[key] = newEle
	if c.list.Len() > c.Capacity {
		delEle := c.list.Back()
		c.list.Remove(delEle)
		delete(c.cacheMap, delEle.Value.(*CacheEntry).key)
	}

}
func (c *LRUCache) Get(key int) int {
	if ele, ok := c.cacheMap[key]; ok {
		c.list.MoveToFront(ele)
		return ele.Value.(*CacheEntry).value
	}
	return -1
}

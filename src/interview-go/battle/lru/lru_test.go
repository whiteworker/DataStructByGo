package lru

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const (
	Expired = 5 * time.Second
)

func TestLruCache(t *testing.T) {
	evictCounter := 0
	onEvicted := func(k Key, v interface{}) {
		evictCounter++
	}
	l := New(16, Expired, onEvicted)

	for i := 0; i < 32; i++ {
		l.Put(i, i, Expired)
	}
	if l.Len() != 16 {
		t.Fatalf("bad len: %v", l.Len())
	}

	if evictCounter != 16 {
		t.Fatalf("bad evict count: %v", evictCounter)
	}

	for i, k := range l.Keys() {
		if v, ok := l.Get(k); !ok || v != k || v != i+16 {
			t.Fatalf("bad key: %v", k)
		}
	}
	for i := 0; i < 16; i++ {
		_, ok := l.Get(i)
		if ok {
			t.Fatalf("should be evicted")
		}
	}
	for i := 16; i < 32; i++ {
		_, ok := l.Get(i)
		if !ok {
			t.Fatalf("should not be evicted")
		}
	}
	for i := 16; i < 24; i++ {
		ok := l.Remove(i)
		if !ok {
			t.Fatalf("should be contained")
		}
		ok = l.Remove(i)
		if ok {
			t.Fatalf("should not be contained")
		}
		_, ok = l.Get(i)
		if ok {
			t.Fatalf("should be deleted")
		}
	}

	l.Get(24)

	l.Clear()
	if l.Len() != 0 {
		t.Fatalf("bad len: %v", l.Len())
	}
	if _, ok := l.Get(30); ok {
		t.Fatalf("should contain nothing")
	}
}

// Test that put returns true/false if an eviction occurred
func TestLRU_Put(t *testing.T) {
	evictCounter := 0
	onEvicted := func(k Key, v interface{}) {
		evictCounter++
		t.Logf("evict k:%v,v:%v", k, v)
	}

	l := New(1, Expired, onEvicted)
	l.Put(1, 1, Expired)
	if evictCounter != 0 {
		t.Errorf("should not have an eviction")
	}
	l.Put(2, 2, Expired)
	if evictCounter != 1 {
		t.Errorf("should have an eviction")
	}
}

// Test that Contains doesn't update recent-ness
func TestLRU_Contains(t *testing.T) {
	l := New(1, Expired, nil)

	l.Put(1, 1, Expired)
	l.Put(2, 2, Expired)
	if !l.Contains(2) {
		t.Errorf("2 should be contained")
	}
	l.Put(3, 3, Expired)
	if l.Contains(1) {
		t.Errorf("Contains should not have updated recent-ness of 1")
	}
}
func TestConcurrent(t *testing.T) {
	l := New(16, Expired, nil)
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(wg *sync.WaitGroup, index int) {
			defer wg.Done()
			l.Put(index, index, Expired)
		}(&wg, i)
	}
	wg.Wait()
	for _, k := range l.Keys() {
		if v, ok := l.Get(k); ok {
			fmt.Printf("key:%v,value:%v \n", k, v)
		} else {
			t.Errorf("could not get value:%v", k)
		}

	}

}

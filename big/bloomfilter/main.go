package main

import (
	"fmt"
	"hash"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	k        uint
	n        uint
	m        uint
	bitset   []bool
	hashFunc []hash.Hash64
}

func NewBloomFilter(size uint, k uint) *BloomFilter {
	var hashFunc []hash.Hash64
	for i := 0; i < int(k); i++ {
		hashFunc = append(hashFunc, murmur3.New64WithSeed(uint32(i)))
	}
	return &BloomFilter{
		bitset:   make([]bool, size),
		k:        k,
		n:        0,
		m:        size,
		hashFunc: hashFunc,
	}
}
func (bf *BloomFilter) AddItem(item []byte) {
	for _, hashFunc := range bf.hashFunc {
		_, err := hashFunc.Write(item)
		if err != nil {
			panic(err)
		}
		hashVal := hashFunc.Sum64()
		position := uint(hashVal) % bf.m
		bf.bitset[position] = true
		hashFunc.Reset()
	}
	bf.n++
}
func (bf *BloomFilter) Verify(item []byte) bool {
	for _, hashFunc := range bf.hashFunc {
		_, err := hashFunc.Write(item)
		if err != nil {
			panic(err)
		}
		hashVal := hashFunc.Sum64()
		pos := uint(hashVal) % bf.m
		if bf.bitset[pos] == false {
			return false
		}
		hashFunc.Reset()
	}
	return true
}
func main() {
	bf := NewBloomFilter(10, 1)
	bf.AddItem([]byte{1})
	fmt.Println(bf.Verify([]byte{1}))
	fmt.Println(bf.Verify([]byte{2}))
}

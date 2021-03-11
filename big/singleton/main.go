package main

import "sync"

type singleton struct{}

var ins *singleton

func GetIns() *singleton {
	sync.Once(func() {
		ins = &singleton{}
	})
	return ins
}

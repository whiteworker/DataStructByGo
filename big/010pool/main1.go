package main

import (
	"sync"
	"time"
)

type Pool struct {
	capacity       int
	running        int
	expiryDuration time.Duration
	workers        []*Worker
	lock           sync.Mutex
	release        chan bool
}

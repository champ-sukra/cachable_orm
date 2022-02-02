package main

import (
	"reflect"
	"sync"
	"time"
)

type Item struct {
	Val interface{}
	Exp int64
}

type CachedMap struct {
	m  map[string]*Item
	mu *sync.RWMutex

	cleaningTicker     *time.Ticker
	cleaningTickerStop chan bool
}

func NewCachedMap() *CachedMap {
	m := &CachedMap{
		m:                  map[string]*Item{},
		mu:                 &sync.RWMutex{},
		cleaningTicker:     time.NewTicker(time.Second * 5),
		cleaningTickerStop: make(chan bool),
	}
	m.StartCleanup()
	return m
}

func (cm *CachedMap) Set(key string, val interface{}, d time.Duration) {
	cm.mu.Lock()
	cm.m[key] = &Item{
		Val: val,
		Exp: time.Now().Add(d).Unix(),
	}
	cm.mu.Unlock()
}

func (cm *CachedMap) Get(m interface{}, key string) bool {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	if item, isFound := cm.m[key]; isFound {
		source := reflect.ValueOf(m)
		dest := reflect.ValueOf(item.Val)
		source.Elem().Set(dest.Elem())
		return true
	}
	return false
}

func (cm *CachedMap) IsExists(key string) bool {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	if _, isFound := cm.m[key]; isFound {
		return true
	}
	return false
}

func (cm *CachedMap) Remove(key string) bool {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if _, isFound := cm.m[key]; isFound {
		delete(cm.m, key)
		return true
	}
	return false
}

func (cm *CachedMap) StartCleanup() {
	go func() {
		for {
			select {
			case <-cm.cleaningTicker.C:
				cm.Clean()
			case <-cm.cleaningTickerStop:
				break
			}
		}
	}()
}

func (cm *CachedMap) Clean() {
	now := time.Now().Unix()
	cm.mu.Lock()

	for key, el := range cm.m {
		if now >= el.Exp {
			delete(cm.m, key)
		}
	}
	cm.mu.Unlock()
}

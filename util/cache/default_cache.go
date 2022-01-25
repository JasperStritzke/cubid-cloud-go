package cache

import (
	"reflect"
	"sync"
	"time"
)

type entry struct {
	value            interface{}
	lastTimeOfAccess time.Time
}

type DefaultCache struct {
	lock        sync.RWMutex
	quitChannel chan bool
	entries     map[interface{}]entry
	timeout     time.Duration
	ticker      *time.Ticker
}

func (d *DefaultCache) Get(key, value interface{}) bool {
	result, exists := d.entries[key]

	if exists {
		d.lock.RLock()
		defer d.lock.RUnlock()

		reflect.ValueOf(value).Elem().Set(reflect.ValueOf(result.value))
		result.lastTimeOfAccess = time.Now()
		return true
	}

	return false
}

func (d *DefaultCache) Exists(key interface{}) bool {
	return d.resetTime(key)
}

func (d *DefaultCache) Put(key, value interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.entries[key] = entry{
		value:            value,
		lastTimeOfAccess: time.Now(),
	}
}

func (d *DefaultCache) Invalidate(key interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()

	delete(d.entries, key)
}

func (d *DefaultCache) resetTime(key interface{}) bool {
	result, exists := d.entries[key]

	if exists {
		d.lock.RLock()
		defer d.lock.RUnlock()

		result.lastTimeOfAccess = time.Now()
		d.entries[key] = result
		return true
	}

	return false
}

func (d *DefaultCache) Shutdown() {
	d.ticker.Stop()
	d.quitChannel <- true
}

func (d *DefaultCache) flushCache() {
	for key, value := range d.entries {
		if time.Now().Sub(value.lastTimeOfAccess) >= d.timeout {
			delete(d.entries, key)

		}
	}
}

func NewCache(timeout time.Duration) *DefaultCache {
	ticker := time.NewTicker(1 * time.Second)

	cache := DefaultCache{
		quitChannel: make(chan bool),
		entries:     make(map[interface{}]entry),
		timeout:     timeout,
		ticker:      ticker,
	}

	go func() {
		for {
			select {
			case <-cache.quitChannel:
				return
			case <-ticker.C:
				cache.flushCache()
			}
		}
	}()

	return &cache
}

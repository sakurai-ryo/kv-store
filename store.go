package main

import (
	"github.com/cornelk/hashmap"
	"sync"
)

type InMemoryStore struct {
	sync.RWMutex
	m *hashmap.HashMap
}

var Store = InMemoryStore{
	m: &hashmap.HashMap{},
}

func Put(key string, value string) error {
	Store.Lock()
	defer Store.Unlock()

	Store.m.Set(key, value)

	return nil
}

func Get(key string) (string, error) {
	Store.RLock()
	defer Store.RUnlock()

	value, ok := Store.m.Get(key)

	if !ok {
		return "", ErrorNoSuchKey
	}

	v := value.(string)

	return v, nil
}

func Delete(key string) error {
	Store.Lock()
	defer Store.Unlock()

	Store.m.Del(key)

	return nil
}

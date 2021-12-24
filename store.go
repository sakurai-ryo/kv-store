package main

import "sync"

type InMemoryStore struct {
	sync.RWMutex
	m map[string]string
}

var Store = InMemoryStore{
	m: make(map[string]string),
}

func Put(key string, value string) error {
	Store.Lock()
	defer Store.Unlock()

	Store.m[key] = value

	return nil
}

func Get(key string) (string, error) {
	Store.RLock()
	defer Store.RUnlock()

	value, ok := Store.m[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	Store.Lock()
	defer Store.Unlock()

	delete(Store.m, key)

	return nil
}

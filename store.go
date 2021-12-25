package main

import (
	"github.com/cornelk/hashmap"
)

type InMemoryStore struct {
	m *hashmap.HashMap
}

var Store = InMemoryStore{
	m: &hashmap.HashMap{},
}

func Put(key string, value string) error {
	Store.m.Set(key, value)

	return nil
}

func Get(key string) (string, error) {
	value, ok := Store.m.Get(key)
	if !ok {
		return "", ErrorNoSuchKey
	}
	v := value.(string)

	return v, nil
}

func Delete(key string) error {
	Store.m.Del(key)

	return nil
}

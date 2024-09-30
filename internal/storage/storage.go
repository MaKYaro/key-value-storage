package storage

import "errors"

var ErrNoSuchKey = errors.New("no such key")

var store = make(map[string]string, 0)

func Put(key string, value string) error {
	store[key] = value

	return nil
}

func Get(key string) (error, string) {
	value, ok := store[key]
	if !ok {
		return ErrNoSuchKey, ""
	}

	return nil, value
}

func Delete(key string) {
	delete(store, key)
}

package storage

import "errors"

type storage struct {
	data map[interface{}]interface{}
}

var notFound = errors.New("key doesn't exist")

func New() storage {
	return storage{make(map[interface{}]interface{})}
}

func (store storage) Get(key interface{}) (interface{}, error) {
	val, found := store.data[key]

	if !found {
		return nil, notFound
	}

	return val, nil
}

func (store storage) GetAll() map[interface{}]interface{} {
	return store.data
}

func (store storage) Put(key interface{}, value interface{}) {
	store.data[key] = value
}

func (store storage) Delete(key interface{}) {
	delete(store.data, key)
}

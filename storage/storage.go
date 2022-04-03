// Package storage implements a type, that contains key-value map
//
// In order to safely manipulate this dictionary, you should use
// the methods of a storage type, and not change the key-value collection directly.
//
// Use the New function to create a storage variable.
//
// Methods Get, GetAll, Put and Delete do exactly what they are called.
// More information about them, you can find in documentation.
package storage

import "errors"

// Type storage is a key-value map with methods that allows you to concurrently
// change data, and do it safely.
//
// -Concurrency will be implemented in sub-task4-
type storage struct {
	data map[interface{}]interface{}
}

var notFound = errors.New("key doesn't exist")

// New returns an empty storage variable, for which memory is allocated for the data.
func New() storage {
	return storage{make(map[interface{}]interface{})}
}

// Get method returns a value from data by key, and also any error encountered.
func (store storage) Get(key interface{}) (interface{}, error) {
	val, found := store.data[key]

	if !found {
		return nil, notFound
	}

	return val, nil
}

// GetAll method returns the whole key-value collection.
func (store storage) GetAll() map[interface{}]interface{} {
	return store.data
}

// Put method creates or rewrites key-value pair in storage.
func (store storage) Put(key interface{}, value interface{}) {
	store.data[key] = value
}

// Delete method annihilates a key-value pair, according to the key, that it receives.
func (store storage) Delete(key interface{}) {
	delete(store.data, key)
}

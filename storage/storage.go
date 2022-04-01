package storage

func Get(id int) string {
	return mainStore.get(id)
}

func Put(value string) int {
	return mainStore.put(value)
}

func Delete(id int) {
	mainStore.delete(id)
}

type storage struct {
	innerRepository map[int]string
	currentId       int
}

var mainStore = storage{make(map[int]string), 0}

func (store storage) get(id int) string {
	return store.innerRepository[id]
}

func (store storage) put(value string) int {
	id := store.currentId
	store.currentId++

	store.innerRepository[id] = value
	return id
}

func (store storage) delete(id int) {
	delete(store.innerRepository, id)
}

package stores

import "sync"

var container = make(map[string]interface{})
var mutex = &sync.Mutex{}

func Set(key string, value interface{}) {
	mutex.Lock()
	container[key] = value
	mutex.Unlock()
}

func Get(key string) interface{} {
	return container[key]
}

func Remove(key string) {
	mutex.Unlock()
	delete(container, key)
	mutex.Unlock()
}

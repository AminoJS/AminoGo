package stores

var container = make(map[string]interface{})

//var mutex = &sync.Mutex{}

func Set(key string, value interface{}) {
	container[key] = value
}

func Get(key string) interface{} {
	return container[key]
}

func Remove(key string) {
	delete(container, key)
}

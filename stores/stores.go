package stores

var container = make(map[string]interface{})

func Set(key string, value interface{}) {
	container[key] = value
}

func Get(key string) interface{} {
	return container[key]
}

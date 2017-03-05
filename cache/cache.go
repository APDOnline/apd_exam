package cache

type APDCache struct {
	Items map[string]interface{}
}

var Cache APDCache

func (*APDCache) Get(key string) (interface{}, bool) {
	if len(Cache.Items) == 0 {
		return nil, false
	}

	if value, ok := Cache.Items[key]; ok {
		return value, true
	} else {
		return nil, false
	}
}

func (*APDCache) Set(key string, value interface{}) *APDCache {
	if len(Cache.Items) == 0 {
		Cache.Items = map[string]interface{}{
			key: value,
		}
	} else {
		Cache.Items[key] = value
	}

	return &Cache
}

func (*APDCache) Remove(key string) {
	if len(Cache.Items) == 0 {
		return
	}

	if _, ok := Cache.Items[key]; ok {
		delete(Cache.Items, key)
	}
}

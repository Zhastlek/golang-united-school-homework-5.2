package cache

import "time"

type Cache struct {
	Hash map[string]user
}

type user struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{
		Hash: make(map[string]user),
	}
}

func (c Cache) Get(key string) (string, bool) {
	if val, ok := c.Hash[key]; ok {
		if val.deadline.IsZero() || time.Now().Before(val.deadline) {
			return val.value, true
		}
	}
	return "", false
}

func (c Cache) Put(key, value string) {
	c.Hash[key] = user{
		value:    value,
		deadline: time.Time{},
	}
}

func (c Cache) Keys() []string {
	var sessions []string
	for key, value := range c.Hash {
		if value.deadline.IsZero() || time.Now().Before(value.deadline) {
			sessions = append(sessions, key)
		}
	}
	return sessions
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.Hash[key] = user{
		value:    value,
		deadline: deadline,
	}
}

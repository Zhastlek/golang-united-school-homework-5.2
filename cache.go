package cache

import "time"

type Cache struct {
	Hash     map[string]string
	Deadline map[string]time.Time
	Exp      map[string]bool
}

func NewCache() Cache {
	return Cache{
		Hash:     make(map[string]string),
		Deadline: make(map[string]time.Time),
		Exp:      make(map[string]bool),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	if value, ok := c.Hash[key]; ok {
		if c.checkExp(key) {
			return value, true
		}
		if c.checkDeadline(key) {
			return value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.Hash[key] = value
	c.Exp[key] = true
	if _, ok := c.Deadline[key]; ok {
		delete(c.Deadline, key)
	}
}

func (c *Cache) Keys() []string {
	var sessions []string
	for key, value := range c.Hash {
		if c.checkExp(key) {
			sessions = append(sessions, value)
			continue
		}
		if c.checkDeadline(key) {
			sessions = append(sessions, value)
			continue
		}
	}
	return sessions
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.Hash[key] = value
	c.Deadline[key] = deadline
	c.Exp[key] = false
}

func (c *Cache) checkDeadline(key string) bool {
	today := time.Now()
	if value, ok := c.Deadline[key]; ok {
		if today.After(value) || today == value {
			return true
		}
	}
	return false
}

func (c *Cache) checkExp(key string) bool {
	return c.Exp[key]
}

func (c *Cache) deleteKey(key string) {
	if _, ok := c.Hash[key]; ok {
		delete(c.Exp, key)
		delete(c.Deadline, key)
		delete(c.Hash, key)
	}
}

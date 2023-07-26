package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	//Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	pnl, ok := c.items[key]
	if ok {
		pnl.Value = value
		c.queue.MoveToFront(pnl)
	} else {
		pnl := c.queue.PushFront(value)
		c.items[key] = pnl
		if c.queue.Len() > c.capacity {
			last := c.queue.Back()
			c.queue.Remove(last)
			// найдем ключ, в котром лежит указатель на послежний элемент очереди. А как иначе?
			// в цикле. а как иначе?
			for k, v := range c.items {
				if v != last {
					continue
				}
				delete(c.items, k)
				break
			}

		}
	}
	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	pnl, ok := c.items[key]
	if ok {
		return pnl.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	// надеюсь, сборщик сделает своё дело ))
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)

}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

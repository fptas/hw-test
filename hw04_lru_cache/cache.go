package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	// Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	pnl, ok := c.items[key] // получить указатель на значение заданного ключа
	if ok { // если в пате был указатель
		pnl.Value = value  // обновим значение в элементе списка
		c.queue.MoveToFront(pnl) // и переведем элемент в начало списка
	} else { // если не было в спарвочнике
		pnl := c.queue.PushFront(value) // добавим в начало
		c.items[key] = pnl // добавим ключ с справочник
		if c.queue.Len() > c.capacity { // если длина списка превышена
			last := c.queue.Back()
			c.queue.Remove(last) // удалим послеждний элемент
			// найдем ключ, в котром лежит указатель на послежний элемент очереди.
			// в цикле. а как иначе?
			for k, v := range c.items {
				if v != last {
					continue
				}
				delete(c.items, k) // удалим ключ из справочника для удаляемого элемента списка
				break
			}
		}
	}
	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	pnl, ok := c.items[key] // попытаемся получит значение спраовчника
	if ok { // если было, вернем
		return pnl.Value, true
	}
	return nil, false // если не было, возвратим ничто
}

func (c *lruCache) Clear() {
	// надеюсь, сборщик сделает своё дело ))
	// заменим объекты новыми
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	// инициализация
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

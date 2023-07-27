package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	// List // Remove me after realization.
	// Place your code here.
	front  *ListItem
	back   *ListItem
	length int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	nl := ListItem{Value: v}
	pnl := &nl
	if l.front == nil {
		l.front = pnl
		l.back = l.front
	} else {
		pnl.Next = l.front
		l.front.Prev = pnl
		l.front = pnl
	}
	l.length++
	return pnl
}

func (l *list) PushBack(v interface{}) *ListItem {
	nl := ListItem{Value: v}
	pnl := &nl
	if l.back == nil {
		l.front = pnl
		l.back = pnl //nolint:gofmt
	} else {
		pnl.Prev = l.back
		l.back.Next = pnl
		l.back = pnl
	}
	l.length++
	return pnl
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front == i {
		return
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	i.Next = l.front
	i.Prev = nil
	l.front.Prev = i
	l.front = i
}

func NewList() List {
	return new(list)
}

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
	nl := ListItem{Value: v}  // создадим новый элемент для списка
	pnl := &nl  // указатель на новый элемент
	if l.front == nil {  // если список пустой, наччало и конец буду смотреть на единственный новый элемент
		l.front = pnl
		l.back = l.front
	} else {
		pnl.Next = l.front // иначе новый элемен тсмотрит на прежний первый
		l.front.Prev = pnl // прежний первый сотрит назад на новый
		l.front = pnl  // начало смотрит на новый элемент
	}
	l.length++ // длина списка увеличилась
	return pnl
}

func (l *list) PushBack(v interface{}) *ListItem {
	nl := ListItem{Value: v} // создадим новый элемент для списка
	pnl := &nl// указатель на новый элемент
	if l.back == nil {   // если список пустой, наччало и конец буду смотреть на единственный новый элемент
		l.front = pnl
		l.back = pnl	//nolint:all //abra-kadabra
	} else {
		pnl.Prev = l.back   // иначе новый элемент смотрит на прежний последний
		l.back.Next = pnl // последний вперед смотрит на новый
		l.back = pnl   // конец смотрит на новый
	}
	l.length++  // длина списка увеличилась
	return pnl
}

func (l *list) Remove(i *ListItem) {
	// соседние элемены удаляемого, если они есть, должны смотреть туда, куда смотрел удаляемый элемент
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	l.length-- // длина списка уменьшилась
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front == i { // если заданный элемент уже первый, то ничего делать не надо
		return
	}
	if i.Next != nil { // как при удалении
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	i.Next = l.front   // как при добавлении вперед
	i.Prev = nil   // но только новый первый элемент назад никуда не смотрит, т.к. он тепепь первый
	l.front.Prev = i
	l.front = i
}

func NewList() List {
	return new(list)
}

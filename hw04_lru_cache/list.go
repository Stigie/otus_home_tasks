package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый Item
	Back() *listItem                   // последний Item
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	mux   sync.Mutex
	Value interface{} // значение
	Next  *listItem   // следующий элемент
	Prev  *listItem   // предыдущий элемент
}

type list struct {
	length       int
	firstElement *listItem
	lastElement  *listItem
	mux          sync.Mutex
	// Place your code here
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *listItem {
	l.mux.Lock()
	defer l.mux.Unlock()
	return l.firstElement
}

func (l *list) Back() *listItem {
	l.mux.Lock()
	defer l.mux.Unlock()
	return l.lastElement
}

func (l *list) PushFront(v interface{}) *listItem {
	l.mux.Lock()
	defer l.mux.Unlock()
	i := listItem{
		Value: v,
		Next:  nil,
		Prev:  l.firstElement,
	}
	if l.firstElement != nil {
		l.firstElement.Next = &i
	} else {
		l.lastElement = &i
	}

	l.firstElement = &i
	l.length++

	return l.firstElement
}

func (l *list) PushBack(v interface{}) *listItem {
	l.mux.Lock()
	defer l.mux.Unlock()
	i := listItem{
		Value: v,
		Next:  l.lastElement,
		Prev:  nil,
	}
	if l.lastElement != nil {
		l.lastElement.Prev = &i
	} else {
		l.firstElement = &i
	}
	l.lastElement = &i
	l.length++
	return l.lastElement
}

func (l *list) Remove(i *listItem) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		i.Next.Prev = nil
		l.lastElement = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		i.Prev.Next = nil
		l.firstElement = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *listItem) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.firstElement == i {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.lastElement = i.Next
		i.Next.Prev = nil
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	i.Prev = l.firstElement
	i.Next = nil
	l.firstElement.Next = i
	l.firstElement = i
}

func NewList() List {
	return &list{}
}

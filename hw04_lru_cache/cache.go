package hw04_lru_cache //nolint:golint,stylecheck

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	Capacity int
	Queue    List
	Items    sync.Map
}

func (lC *lruCache) CleanCacheIfFull() {
	if lC.Capacity == lC.Queue.Len() {
		lastElem := lC.Queue.Back()
		lastElem.mux.Lock()
		defer lastElem.mux.Unlock()
		lC.Items.Delete(lastElem.Value.(*cacheItem).Key)
		lC.Queue.Remove(lastElem)
	}
}

func (lC *lruCache) Set(key Key, value interface{}) bool {
	storeValue, ok := lC.Items.Load(key)
	if ok {
		itemQuery := storeValue.(*cacheItem)
		itemQuery.mux.Lock()
		defer itemQuery.mux.Unlock()
		itemQuery.Value = value
		itemQuery.QueueLink.Value = itemQuery
		lC.Queue.MoveToFront(itemQuery.QueueLink)
	} else {
		lC.CleanCacheIfFull()
		item := lC.Queue.PushFront(&cacheItem{
			Key:       key,
			Value:     value,
			QueueLink: nil,
		})
		item.Value.(*cacheItem).QueueLink = item

		lC.Items.Store(key, &cacheItem{
			Key:       key,
			Value:     value,
			QueueLink: item,
		})
	}

	return ok
}

func (lC *lruCache) Get(key Key) (interface{}, bool) {
	storeValue, ok := lC.Items.Load(key)
	if ok {
		itemQuery := storeValue.(*cacheItem)
		itemQuery.mux.Lock()
		defer itemQuery.mux.Unlock()
		itemQuery.QueueLink.Value = itemQuery
		lC.Queue.MoveToFront(itemQuery.QueueLink)
		return storeValue.(*cacheItem).Value, ok
	}
	return nil, false
}
func (lC *lruCache) Clear() {
	lC.Items = sync.Map{}
	lC.Queue = NewList()
}

type cacheItem struct {
	Key
	Value     interface{}
	QueueLink *listItem
	mux sync.Mutex
}

func NewCache(capacity int) Cache {
	queue := NewList()
	return &lruCache{
		Capacity: capacity,
		Queue:    queue,
	}
}

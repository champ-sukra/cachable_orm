package orm_handler

import (
	"sync/atomic"
)

type UidGenerator struct{}

type Item struct {
	K string
	V uint32
}

var atmItems []*Item

func NewUidGenerator() *UidGenerator {
	return &UidGenerator{}
}

func (c *UidGenerator) SetUniqueIdForKey(key string, val uint32) {
	var item Item
	isFound := c.Find(&item, key)
	if !isFound {
		item.K = key
		item.V = val
		atomic.StoreUint32(&item.V, val)
		atmItems = append(atmItems, &item)
	}
}

func (c *UidGenerator) GetUniqueIdForKey(key string) uint32 {
	var item *Item
	for _, ai := range atmItems {
		if ai.K == key {
			item = ai
			break
		}
	}
	return atomic.AddUint32(&item.V, 1)
}

func (c *UidGenerator) Find(item *Item, key string) bool {
	for _, ai := range atmItems {
		if ai.K == key {
			item = ai
			return true
		}
	}
	return false
}

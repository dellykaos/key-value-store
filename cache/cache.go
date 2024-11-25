package cache

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
)

var (
	errDataType = errors.New("Data Type Error")
)

type Cache struct {
	mu             sync.Mutex
	cache          map[string]*node
	attributesType map[string]valueType
	linkedList     *linkedList
}

func NewCache() *Cache {
	return &Cache{
		cache:          map[string]*node{},
		attributesType: map[string]valueType{},
		linkedList:     newLinkedList(),
	}
}

func (c *Cache) Get(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, found := c.cache[key]
	if !found {
		return "", fmt.Errorf("No entry found for %s", key)
	}

	return node.string(), nil
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, found := c.cache[key]
	if !found {
		return
	}

	c.linkedList.delete(node)
	delete(c.cache, key)
}

func (c *Cache) Put(key string, attributeKeys, attributeValues []string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, found := c.cache[key]
	if found {
		items, err := c.buildItems(attributeKeys, attributeValues)
		if err != nil {
			return err
		}

		node.attributes = items
		return nil
	}

	items, err := c.buildItems(attributeKeys, attributeValues)
	if err != nil {
		return err
	}

	node = newNode(key, items)
	c.linkedList.add(node)
	c.cache[key] = node
	return nil
}

func (c *Cache) Search(attrKey, attrVal string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	res := []string{}
	n := c.linkedList.head
	for n != c.linkedList.tail {
		if n.attributes == nil {
			n = n.next
			continue
		}

		for _, attr := range n.attributes {
			if attr.key == attrKey && attr.value == attrVal {
				res = append(res, n.key)
				break
			}
		}
		n = n.next
	}

	sort.Strings(res)
	return strings.Join(res, ",")
}

func (c *Cache) Keys() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys := []string{}
	n := c.linkedList.head
	for n != c.linkedList.tail {
		if n.attributes == nil {
			n = n.next
			continue
		}
		keys = append(keys, n.key)
		n = n.next
	}

	sort.Strings(keys)
	return strings.Join(keys, ",")
}

func (c *Cache) buildItems(attributeKeys, attributeValues []string) ([]*item, error) {
	attributes := []*item{}
	for i := 0; i < len(attributeKeys); i++ {
		k, v := attributeKeys[i], attributeValues[i]
		item := newItem(k, v)
		if t, found := c.attributesType[k]; found {
			if t != item.valType {
				return nil, errDataType
			}
		} else {
			item = newItem(k, v)
			c.attributesType[k] = item.valType
		}
		attributes = append(attributes, item)
	}
	return attributes, nil
}

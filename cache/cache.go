package cache

import (
	"errors"
	"fmt"
)

var (
	errDataType = errors.New("Data Type Error")
)

type Cache struct {
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
	node, found := c.cache[key]
	if !found {
		return "", fmt.Errorf("No entry found for %s", key)
	}

	return node.String(), nil
}

func (c *Cache) Delete(key string) {
	node, found := c.cache[key]
	if !found {
		return
	}

	c.linkedList.delete(node)
	delete(c.cache, key)
}

func (c *Cache) Put(key string, values map[string]string) error {
	node, found := c.cache[key]
	if found {
		items, err := c.buildItems(values)
		if err != nil {
			return err
		}

		node.attributes = items
		return nil
	}

	items, err := c.buildItems(values)
	if err != nil {
		return err
	}

	node = newNode(key, items)
	c.linkedList.add(node)
	c.cache[key] = node
	return nil
}

func (c *Cache) Search(attrKey, attrVal string) string {
	res := ""
	n := c.linkedList.head
	for n != c.linkedList.tail {
		if n.attributes == nil {
			n = n.next
			continue
		}

		for _, attr := range n.attributes {
			if attr.key == attrKey && attr.value == attrVal {
				res += fmt.Sprintf("%s,", n.key)
				break
			}
		}
		n = n.next
	}

	return res[0 : len(res)-1]
}

func (c *Cache) Keys() string {
	keys := ""
	n := c.linkedList.head
	for n != c.linkedList.tail {
		if n.attributes == nil {
			n = n.next
			continue
		}
		keys += n.key + ","
		n = n.next
	}

	return keys[0 : len(keys)-1]
}

func (c *Cache) buildItems(values map[string]string) ([]*item, error) {
	attributes := []*item{}
	for k, v := range values {
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

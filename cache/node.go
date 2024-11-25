package cache

type node struct {
	key        string
	attributes []*item
	prev       *node
	next       *node
}

func newNode(key string, attributes []*item) *node {
	return &node{
		key:        key,
		attributes: attributes,
	}
}

func (n *node) string() string {
	res := ""
	for i, item := range n.attributes {
		res += item.string()
		if i < len(n.attributes)-1 {
			res += ", "
		}
	}
	return res
}

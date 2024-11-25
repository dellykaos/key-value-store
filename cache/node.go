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

func (n *node) String() string {
	res := ""
	for i, item := range n.attributes {
		res += item.String()
		if i < len(n.attributes)-1 {
			res += ", "
		}
	}
	return res
}

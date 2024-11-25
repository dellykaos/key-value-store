package cache

import "fmt"

type valueType string

const (
	StringValueType  valueType = "string"
	IntegerValueType valueType = "integer"
	DoubleValueType  valueType = "double"
	BooleanValueType valueType = "boolean"
)

type item struct {
	key     string
	value   string
	valType valueType
}

func newItem(key, value string) *item {
	item := &item{
		key:   key,
		value: value,
	}
	switch {
	case isFloat(value):
		item.valType = DoubleValueType
	case isInt(value):
		item.valType = IntegerValueType
	case isBool(value):
		item.valType = BooleanValueType
	default:
		item.valType = StringValueType
	}

	return item
}

func (item *item) String() string {
	return fmt.Sprintf("%s: %s", item.key, item.value)
}

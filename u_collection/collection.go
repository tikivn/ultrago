package collection

func NewStructCollection(items ...interface{}) *StructCollection {
	return &StructCollection{items: items}
}

type StructCollection struct {
	items []interface{}
}

func (c *StructCollection) Find(f func(item interface{}) bool) interface{} {
	for _, item := range c.items {
		found := f(item)
		if found {
			return item
		}
	}
	return nil
}

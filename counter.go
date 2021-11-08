package counter

import "reflect"

// Counter is used to count hashable items,
// Sometimes called a bag or multiset.
// Elements are stored as dictionary keys and their counts
// are stored as dictionary values.
type Counter struct {
	ori  interface{}
	data map[interface{}]interface{}
}

func NewNullCounter() *Counter {
	return &Counter{}
}

func NewCounter(item interface{}) *Counter {
	// item:struct,string,map,slice
	switch t := reflect.TypeOf(item); t.Name() {
	case "string":
	case "struct":
	default:

	}
	c := Counter{}

	return &c
}

func (c *Counter) Elements() {

}

func (c *Counter) Items() *map[interface{}]interface{} {
	return &c.data
}

func (c *Counter) Pop() {

}

func (c *Counter) PopItem() {

}

func (c *Counter) Keys() {

}
func (c *Counter) Values() {

}

func (c *Counter) Get() {

}

func (c *Counter) Clear() {

}

func (c *Counter) Copy() {

}

func (c *Counter) FromKeys() {

}

func (c *Counter) Update() {

}

func (c *Counter) MostCommon() {

}

func (c *Counter) Subtract() {

}

func (c *Counter) SetDefault() {

}

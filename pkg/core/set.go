package core

type (
	Set struct {
		hash map[interface{}]dummy
	}

	dummy struct{}
)

// Create a new set
func NewSet(initial ...interface{}) *Set {
	s := &Set{make(map[interface{}]dummy)}

	for _, v := range initial {
		s.Put(v)
	}

	return s
}

// Add an element to the set
func (this *Set) Put(element interface{}) {
	this.hash[element] = dummy{}
}

// Remove an element from the set
func (this *Set) Delete(element interface{}) {
	delete(this.hash, element)
}

// Return the number of items in the set
func (this *Set) Len() int {
	return len(this.hash)
}

// Test to see whether or not the element is in the set
func (this *Set) Has(element interface{}) bool {
	_, exists := this.hash[element]
	return exists
}

// Call f for each item in the set
func (this *Set) Do(f func(interface{})) {
	for k, _ := range this.hash {
		f(k)
	}
}

// Find the difference between two sets
func (this *Set) Difference(set *Set) *Set {
	n := make(map[interface{}]dummy)

	for k, _ := range this.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = dummy{}
		}
	}

	return &Set{n}
}

// Find the intersection of two sets
func (this *Set) Intersection(set *Set) *Set {
	n := make(map[interface{}]dummy)

	for k, _ := range this.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = dummy{}
		}
	}

	return &Set{n}
}

// Test whether or not this set is a subset of "set"
func (this *Set) Contains(set *Set) bool {
	if this.Len() > set.Len() {
		return false
	}
	for k, _ := range this.hash {
		if _, exists := set.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Union of two sets
func (this *Set) Union(set *Set) *Set {
	n := make(map[interface{}]dummy)

	for k, _ := range this.hash {
		n[k] = dummy{}
	}
	for k, _ := range set.hash {
		n[k] = dummy{}
	}

	return &Set{n}
}

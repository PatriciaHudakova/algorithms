package main

// IntegerSet stores an int and a bool, bool is true if entry is even
type IntegerSet struct {
	entry map[int]bool
}

// Has return true if the element exists in the set
func (set *IntegerSet) Has(item int) bool {
	_, found := set.entry[item]
	if !found {
		return false
	}
	return true
}

// Add adds a new element to the Set
func (set *IntegerSet) Add(item int) *IntegerSet {
	// if the int is prime return true
	decision := false
	if item%2 == 0 {
		decision = true
	}

	// initialise a new map if the entry is nil
	if set.entry == nil {
		set.entry = make(map[int]bool)
	}

	// and/or assign a new entry to it
	set.entry[item] = decision
	return set
}

// Delete removes an entry from set by key
func (set *IntegerSet) Delete(item int) {
	delete(set.entry, item)
}

// Items returns all keys in the set
func (set *IntegerSet) Items() []int {
	var items []int
	for k := range set.entry {
		items = append(items, k)
	}
	return items
}

// Clear removes all elements from set
func (set *IntegerSet) Clear() {
	set.entry = make(map[int]bool)
}

// Size returns the size of set
func (set IntegerSet) Size() int {
	return len(set.entry)
}

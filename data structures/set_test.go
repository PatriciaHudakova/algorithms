package main

import (
	"testing"
)

func TestIntegerSet(t *testing.T) {
	set := IntegerSet{}
	set.Add(4)
	set.Add(5)

	if set.Size() != 2 {
		t.Fatalf("Expected size of set: %v, got: %v", 2, set.Size())
	}
	if set.entry[4] != true {
		t.Fatalf("Expected boolean: %v, got: %v", true, set.entry[4])
	}
	if set.entry[5] != false {
		t.Fatalf("Expected boolean: %v, got: %v", false, set.entry[5])
	}
	if result := set.Has(4); result != true {
		t.Fatalf("Expected boolean: %v, got: %v", true, set.Has(4))
	}
	if result := set.Has(6); result != false {
		t.Fatalf("Expected boolean: %v, got: %v", false, set.Has(4))
	}
	if len(set.Items()) != 2 {
		t.Fatalf("Incorrect size of items, expected %v, got %v", 2, set.Items())
	}
}

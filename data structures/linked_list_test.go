package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Add(t *testing.T) {
	testList := linkedList{}
	testList.Add(&node{"test data 1", nil})
	testList.Add(&node{"test data 2", nil})

	assert.Equal(t, 2, testList.length)
	assert.Equal(t, "test data 1", testList.head.data)
	assert.Equal(t, "test data 2", testList.tail.data)
}

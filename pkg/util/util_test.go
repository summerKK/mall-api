package util_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/summerKK/mall-api/pkg/util"
)

func TestIsSliceElemPtr(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	elem0 := []*Person{
		&Person{
			Name: "summer",
			Age:  28,
		},
	}

	elem1 := &Person{}
	elem2 := "summer"

	assert.Equal(t, true, util.IsSliceElemPtr(elem0))
	assert.Equal(t, false, util.IsSliceElemPtr(elem1))
	assert.Equal(t, false, util.IsSliceElemPtr(elem2))
}

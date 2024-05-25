package main_test

import (
	"reflect"
	"testing"

	"github.com/nemesidaa/gohashlist/v1/pkg/set"
)

func TestIntersection(t *testing.T) {
	setA := &set.Set{Set: map[interface{}]struct{}{"a": {}, "b": {}}}
	setB := &set.Set{Set: map[interface{}]struct{}{"a": {}, "d": {}}}

	expected := &set.Set{Set: map[interface{}]struct{}{"a": {}}}

	result, err := set.Intersection(setA, setB)
	if err != nil {
		t.Fatalf("failed to calculate intersection: %v", err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, found %v", expected.Set, result.Set)
	}
}

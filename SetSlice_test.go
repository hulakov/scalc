package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	a SetSlice
	b SetSlice
	r SetSlice
}

func TestUnion(t *testing.T) {
	cc := []TestCase{
		TestCase{a: []int{1, 2, 3}, b: []int{4, 5, 6}, r: []int{1, 2, 3, 4, 5, 6}},
		TestCase{a: []int{1, 3, 5}, b: []int{2, 4, 6}, r: []int{1, 2, 3, 4, 5, 6}},
		TestCase{a: []int{1, 2, 3}, b: []int{1, 2, 3}, r: []int{1, 2, 3}},
		TestCase{a: []int{1, 2}, b: []int{2, 3}, r: []int{1, 2, 3}},
		TestCase{a: []int{1}, b: []int{}, r: []int{1}},
		TestCase{a: []int{}, b: []int{1}, r: []int{1}},
		TestCase{a: []int{}, b: []int{}, r: []int{}},
	}

	for _, c := range cc {
		v := c.a.Union(c.b)
		if !reflect.DeepEqual(v, c.r) {
			t.Errorf("Expected the union of %v and %v to be %v but instead got %v!", c.a, c.b, c.r, v)
		}
	}
}

func TestIntersection(t *testing.T) {
	cc := []TestCase{
		TestCase{a: []int{1, 2, 3}, b: []int{1}, r: []int{1}},
		TestCase{a: []int{1, 2, 3}, b: []int{2}, r: []int{2}},
		TestCase{a: []int{1, 2, 3}, b: []int{3}, r: []int{3}},
		TestCase{a: []int{1, 2, 3}, b: []int{}, r: []int{}},
		TestCase{a: []int{}, b: []int{1, 2, 3}, r: []int{}},
		TestCase{a: []int{1, 2}, b: []int{2, 3}, r: []int{2}},
	}

	for _, c := range cc {
		v := c.a.Intersection(c.b)
		if !reflect.DeepEqual(v, c.r) {
			t.Errorf("Expected the intersection of %v and %v to be %v but instead got %v!", c.a, c.b, c.r, v)
		}
	}
}

func TestDifference(t *testing.T) {
	cc := []TestCase{
		TestCase{a: []int{1, 2, 3}, b: []int{1}, r: []int{2, 3}},
		TestCase{a: []int{1, 2, 3}, b: []int{2}, r: []int{1, 3}},
		TestCase{a: []int{1, 2, 3}, b: []int{3}, r: []int{1, 2}},
		TestCase{a: []int{1, 2, 3}, b: []int{}, r: []int{1, 2, 3}},
		TestCase{a: []int{1, 2}, b: []int{2, 3}, r: []int{1}},
		TestCase{a: []int{}, b: []int{1, 2, 3}, r: []int{}},
	}

	for _, c := range cc {
		v := c.a.Difference(c.b)
		if !reflect.DeepEqual(v, c.r) {
			t.Errorf("Expected the intersection of %v and %v to be %v but instead got %v!", c.a, c.b, c.r, v)
		}
	}
}

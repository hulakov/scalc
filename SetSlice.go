package main

import (
	"bufio"
	"os"
	"strconv"
)

type SetSlice []int

func ReadSetFromFile(path string) (SetSlice, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := SetSlice{}
	for scanner.Scan() {
		v, e := strconv.Atoi(scanner.Text())
		if e != nil {
			return nil, e
		}
		result = append(result, v)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return result, nil
}

func (a SetSlice) Union(b SetSlice) SetSlice {
	r := SetSlice{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			r = append(r, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			r = append(r, a[i])
			i++
		} else {
			r = append(r, b[j])
			j++
		}
	}
	r = append(r, a[i:]...)
	r = append(r, b[j:]...)
	return r
}

func (a SetSlice) UnionN(bb []SetSlice) SetSlice {
	result := a
	for _, b := range bb {
		result = result.Union(b)
	}
	return result
}

func (a SetSlice) Intersection(b SetSlice) SetSlice {
	r := SetSlice{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			r = append(r, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return r
}

func (a SetSlice) IntersectionN(bb []SetSlice) SetSlice {
	result := a
	for _, b := range bb {
		result = result.Intersection(b)
	}
	return result
}

func (a SetSlice) Difference(b SetSlice) SetSlice {
	r := SetSlice{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else if a[i] < b[j] {
			r = append(r, a[i])
			i++
		} else {
			j++
		}
	}
	r = append(r, a[i:]...)
	return r
}

func (a SetSlice) DifferenceN(bb []SetSlice) SetSlice {
	result := a
	for _, b := range bb {
		result = result.Difference(b)
	}
	return result
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	command := strings.Join(os.Args[1:], " ")

	p := &Parser{}
	p.AddOperation("SUM", func(ss []SetSlice) (SetSlice, error) {
		if len(ss) == 0 {
			return nil, errors.New("SUM expects 1 or more arguments")
		}
		return ss[0].UnionN(ss[1:]), nil
	})
	p.AddOperation("INT", func(ss []SetSlice) (SetSlice, error) {
		if len(ss) == 0 {
			return nil, errors.New("INT expects 1 or more arguments")
		}
		return ss[0].IntersectionN(ss[1:]), nil
	})
	p.AddOperation("DIF", func(ss []SetSlice) (SetSlice, error) {
		if len(ss) == 0 {
			return nil, errors.New("DIF expects 1 or more arguments")
		}
		return ss[0].DifferenceN(ss[1:]), nil
	})
	s, e := p.Parse(command)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(s)
	}

}

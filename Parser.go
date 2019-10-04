package main

import (
	"errors"
	"fmt"
)

type Operation func([]SetSlice) (SetSlice, error)

type Parser struct {
	operations map[string]Operation
}

func (p *Parser) AddOperation(name string, operation Operation) {
	if p.operations == nil {
		p.operations = make(map[string]Operation)
	}
	p.operations[name] = operation
}

func makeError(t *Token, format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	if t == nil {
		return errors.New(message)
	}
	return errors.New(fmt.Sprintf("%v (position %v)", message, t.Position))
}

func (p *Parser) doParse(c <-chan *Token) (SetSlice, error) {
	t, ok := <-c
	if !ok {
		return nil, makeError(nil, "unexpected end of expression")
	}
	if t.Type != LITERAL {
		return nil, makeError(t, "syntax error, expected operation name")
	}
	operation, ok := p.operations[t.Text]
	if !ok {
		return nil, makeError(t, "bad operation name '%v'", t.Text)
	}

	var args []SetSlice
	for {
		t, ok := <-c
		if !ok {
			return nil, makeError(nil, "the expression should new with ]")
		}
		switch t.Type {
		case OPENING_BRACKET:
			s, e := p.doParse(c)
			if e != nil {
				return nil, e
			}
			args = append(args, s)
		case CLOSING_BRACKET:
			return operation(args)
		case LITERAL:
			s, e := ReadSetFromFile(t.Text)
			if e != nil {
				return nil, makeError(t, "cannot read set from file '%v': %v", t.Text, e)
			}
			args = append(args, s)
		}
	}
}

func (p *Parser) Parse(text string) (SetSlice, error) {
	c := Tokenize(text)

	t, ok := <-c
	if !ok {
		return SetSlice{}, nil
	}

	if t.Type != OPENING_BRACKET {
		return nil, makeError(t, "The expression should start with [")
	}

	s, e := p.doParse(c)
	if e != nil {
		return nil, e
	}

	t, ok = <-c
	if ok {
		return nil, makeError(t, "syntax error, extra symbols")
	}

	return s, nil
}

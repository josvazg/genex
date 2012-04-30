package genex

import (
	"fmt"
)

const (
	PrimitivesOnly = iota
	Any
)

type GenericSnippet struct {
	name, desc string
	genType int
	code string
}

type Genex interface {
	list() []GenericSnippet
	get(name string) GenericSnippet
	put(gen GenericSnippet)
	remove(name string)
}

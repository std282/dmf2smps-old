package smpsbuild

import (
	"container/list"
	"io"
)

// fromBytes describes anything that could be transformed to bytes
type fromBytes interface {
	represent(w io.Writer)
	size() uint
}

// Pattern represents any sequence of SMPS events.
// Pattern is SMPS event, too.
type Pattern struct {
	events     *list.List
	references *list.List

	lastNote   byte
	lastLength int8
}

// NewPattern returns new pattern ready to use
func NewPattern() *Pattern {
	pat := new(Pattern)
	pat.events = list.New()
	pat.references = list.New()
	return pat
}

func (pat *Pattern) foreachEvent(action func(fb fromBytes)) {
	for node := pat.events.Front(); node != nil; node = node.Next() {
		fb := node.Value.(fromBytes)
		action(fb)
	}
}

func (pat *Pattern) foreachRef(action func(addr address)) {
	for node := pat.references.Front(); node != nil; node = node.Next() {
		addr := node.Value.(address)
		action(addr)
	}
}

func (pat *Pattern) represent(w io.Writer) {
	pat.foreachEvent(func(fb fromBytes) {
		fb.represent(w)
	})
}

func (pat *Pattern) size() uint {
	var size uint
	pat.foreachEvent(func(fb fromBytes) {
		size += fb.size()
	})

	return size
}

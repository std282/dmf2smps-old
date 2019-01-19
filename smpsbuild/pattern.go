package smpsbuild

import (
	"bytes"
	"container/list"
)

// fromBytes describes anything that could be transformed to bytes
type fromBytes interface {
	represent() []byte
	size() uint
}

// Pattern represents any sequence of SMPS events.
// Pattern is SMPS event, too.
type Pattern struct {
	events     *list.List
	references *list.List
}

// NewPattern returns new pattern ready to use
func NewPattern() *Pattern {
	return &Pattern{
		events:     list.New(),
		references: list.New(),
	}
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

func (pat *Pattern) represent() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, pat.size()))

	pat.foreachEvent(func(fb fromBytes) {
		buf.Write(fb.represent())
	})

	return buf.Bytes()
}

func (pat *Pattern) size() uint {
	var size uint
	pat.foreachEvent(func(fb fromBytes) {
		size += fb.size()
	})

	return size
}

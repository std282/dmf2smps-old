package smpsbuild

/*
	This file describes SMPS pattern - addressable list of SMPS events, its
	constructor and basic helper functions.
*/

import (
	"container/list"
	"io"
)

// fromBytes describes anything that could be transformed to bytes.
type fromBytes interface {
	represent(w io.Writer) // writes binary representation to writer
	size() uint            // returns size in bytes of possible representation
}

// Pattern represents any sequence of SMPS events.
// Pattern is SMPS event, too.
type Pattern struct {
	events     *list.List
	references *list.List

	eventsSize uint
}

// NewPattern returns new pattern ready to use
func NewPattern() *Pattern {
	pat := new(Pattern)
	pat.events = list.New()
	pat.references = list.New()
	return pat
}

// foreachEvent iterates on underlying list of pattern events with specified
// function.
func (pat *Pattern) foreachEvent(action func(fb fromBytes)) {
	for node := pat.events.Front(); node != nil; node = node.Next() {
		fb := node.Value.(fromBytes)
		action(fb)
	}
}

// foreachRef iterates on underlying list of pattern references with specified
// function.
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

// setInnerPointers sets position of each relative pointer found in pattern,
// given pattern position.
//
// It iterates on each event, computes its size, therefore computes its position
// and if event contains relative pointer, sets its position to specified.
func (pat *Pattern) setInnerPointers(pos uint) {
	for node := pat.events.Front(); node != nil; node = node.Next() {
		fb := node.Value.(fromBytes)

		if raEvent, ok := node.Value.(relAddrEvent); ok {
			raEvent.setPosition(pos)
		}

		pos += fb.size()
	}
}

// addEvent adds event to pattern and increases size of pattern accordingly.
func (pat *Pattern) addEvent(fb fromBytes) {
	pat.events.PushBack(fb)
	pat.eventsSize += fb.size()
}

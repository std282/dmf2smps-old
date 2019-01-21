package smpsbuild

/*
	This file contains functions for adding coordination flags for song flow
	control:

	$F6 - jump
	$F7 - loop
	$F8 - call
	$E3 - return from call
	$F2 - stop

	And it also contains interface for setting relative pointer in such events.
*/

import (
	"io"
)

// relAddrEvent describes any event with relative pointer
type relAddrEvent interface {
	setPosition(pos uint) // sets relative pointer position
}

// PlaceJump adds $F6 (jump) coordination flag.
func (pat *Pattern) PlaceJump(destination *Pattern) {
	jump := newEventJump()
	destination.addRef(jump.where)

	pat.addEvent(jump)
}

type eventJump struct {
	where *relativeAddress
}

func newEventJump() *eventJump {
	jump := new(eventJump)
	jump.where = new(relativeAddress)
	return jump
}

func (jump *eventJump) represent(w io.Writer) {
	w.Write([]byte{0xF6})
	jump.where.represent(w)
}

func (*eventJump) size() uint {
	return 3
}

func (jump *eventJump) setPosition(pos uint) {
	jump.where.pointerPosition = pos + 1 // 1 is for flag byte
}

// PlaceCall adds $F8 (call) coordination flag.
func (pat *Pattern) PlaceCall(destination *Pattern) {
	call := newEventCall()
	destination.addRef(call.where)

	pat.addEvent(call)
}

type eventCall struct {
	where *relativeAddress
}

func newEventCall() *eventCall {
	call := new(eventCall)
	call.where = new(relativeAddress)
	return call
}

func (call *eventCall) represent(w io.Writer) {
	w.Write([]byte{0xF8})
	call.where.represent(w)
}

func (*eventCall) size() uint {
	return 3
}

func (call *eventCall) setPosition(pos uint) {
	call.where.pointerPosition = pos + 1 // 1 is for flag byte
}

// PlaceLoop adds $F7 (loop) coordination flag.
//
// It effectively makes this pattern repeat itself several times.
func (pat *Pattern) PlaceLoop(priority, times int) {
	loop := newEventLoop()
	pat.addRef(loop.start)
	loop.priority = byte(priority)
	loop.repeats = uint8(times)

	pat.addEvent(loop)
}

type eventLoop struct {
	start    *relativeAddress
	priority byte
	repeats  uint8
}

func newEventLoop() *eventLoop {
	loop := new(eventLoop)
	loop.start = new(relativeAddress)
	return loop
}

func (loop *eventLoop) represent(w io.Writer) {
	w.Write([]byte{
		0xF7,
		loop.priority,
		byte(loop.repeats),
	})

	loop.start.represent(w)
}

func (*eventLoop) size() uint {
	return 5
}

func (loop *eventLoop) setPosition(pos uint) {
	loop.start.pointerPosition = pos + 3 // 3 is for flag bytes
}

// PlaceStop places $F2 (stop) coordination flag.
func (pat *Pattern) PlaceStop() {
	stop := new(eventStop)

	pat.addEvent(stop)
}

type eventStop struct{}

func (*eventStop) represent(w io.Writer) {
	w.Write([]byte{0xF2})
}

func (*eventStop) size() uint {
	return 1
}

// PlaceReturn adds $E3 (return after call) coordination flag.
func (pat *Pattern) PlaceReturn() {
	ret := new(eventReturn)

	pat.addEvent(ret)
}

type eventReturn struct{}

func (*eventReturn) represent(w io.Writer) {
	w.Write([]byte{0xE3})
}

func (*eventReturn) size() uint {
	return 1
}

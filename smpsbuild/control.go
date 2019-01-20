package smpsbuild

import (
	"io"
)

// PlaceJump adds $F6 (jump) coordination flag.
func (pat *Pattern) PlaceJump(destination *Pattern) {
	jump := newEventJump()
	destination.addRef(jump.Where)

	pat.events.PushBack(jump)
}

type eventJump struct {
	Where *relativeAddress
}

func newEventJump() *eventJump {
	jump := new(eventJump)
	jump.Where = new(relativeAddress)
	return jump
}

func (jump *eventJump) represent(w io.Writer) {
	w.Write([]byte{0xF6})
	jump.Where.represent(w)
}

func (*eventJump) size() uint {
	return 3
}

// PlaceCall adds $F8 (call) coordination flag.
func (pat *Pattern) PlaceCall(destination *Pattern) {
	call := newEventCall()
	destination.addRef(call.Where)

	pat.events.PushBack(call)
}

type eventCall struct {
	Where *relativeAddress
}

func newEventCall() *eventCall {
	call := new(eventCall)
	call.Where = new(relativeAddress)
	return call
}

func (call *eventCall) represent(w io.Writer) {
	w.Write([]byte{0xF8})
	call.Where.represent(w)
}

func (*eventCall) size() uint {
	return 3
}

// PlaceLoop adds $F7 (loop) coordination flag.
//
// It effectively makes this pattern repeat itself several times.
func (pat *Pattern) PlaceLoop(priority, times int) {
	loop := newEventLoop()
	pat.addRef(loop.Start)
	loop.Priority = byte(priority)
	loop.Repeats = uint8(times)

	pat.events.PushBack(loop)
}

type eventLoop struct {
	Start    *relativeAddress
	Priority byte
	Repeats  uint8
}

func newEventLoop() *eventLoop {
	loop := new(eventLoop)
	loop.Start = new(relativeAddress)
	return loop
}

func (loop *eventLoop) represent(w io.Writer) {
	w.Write([]byte{
		0xF7,
		loop.Priority,
		byte(loop.Repeats),
	})

	loop.Start.represent(w)
}

func (*eventLoop) size() uint {
	return 5
}

// PlaceStop places $F2 (stop) coordination flag.
func (pat *Pattern) PlaceStop() {
	stop := new(eventStop)

	pat.events.PushBack(stop)
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

	pat.events.PushBack(ret)
}

type eventReturn struct{}

func (*eventReturn) represent(w io.Writer) {
	w.Write([]byte{0xE3})
}

func (*eventReturn) size() uint {
	return 1
}

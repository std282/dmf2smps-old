package smpsbuild

import "bytes"

type eventJump struct {
	Where relativeAddress
}

func (jump *eventJump) represent() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 3))
	buf.WriteByte(0xF6)
	buf.Write(jump.Where.represent())

	return buf.Bytes()
}

func (*eventJump) size() uint {
	return 3
}

type eventCall struct {
	Where relativeAddress
}

func (call *eventCall) represent() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 3))
	buf.WriteByte(0xF8)
	buf.Write(call.Where.represent())

	return buf.Bytes()
}

func (*eventCall) size() uint {
	return 3
}

type eventLoop struct {
	Start    relativeAddress
	Priority byte
	Repeats  uint8
}

func (loop *eventLoop) represent() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 5))
	buf.Write([]byte{
		0xF7,
		loop.Priority,
		byte(loop.Repeats),
	})

	buf.Write(loop.Start.represent())

	return buf.Bytes()
}

func (*eventLoop) size() uint {
	return 5
}

type eventStop struct{}

func (*eventStop) represent() []byte {
	return []byte{
		0xF2,
	}
}

func (*eventStop) size() uint {
	return 1
}

type eventReturn struct{}

func (*eventReturn) represent() []byte {
	return []byte{0xE3}
}

func (*eventReturn) size() uint {
	return 1
}

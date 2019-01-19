package smpsbuild

import (
	"bytes"
	"encoding/binary"
)

type address interface {
	evaluate(pos uint) // sets position of pointed-to
	represent() []byte // writes address in SMPS notation
	isNull() bool      // returns true if address doesn't point at anything
}

type relativeAddress struct {
	PointerPosition uint // where is pointer located
	EntityPosition  uint // where is pointed-to located
}

func newRelAddr() *relativeAddress {
	return &relativeAddress{
		PointerPosition: 0,
		EntityPosition:  0,
	}
}

func (relAddr *relativeAddress) evaluate(pos uint) {
	if relAddr.EntityPosition != 0 {
		logger.Printf(
			"warning: relative pointer reevaluation (%d -> %d)",
			relAddr.EntityPosition,
			pos,
		)
	}

	relAddr.EntityPosition = pos
}

func (relAddr *relativeAddress) represent() []byte {
	if relAddr.isNull() {
		logger.Fatal(
			"error: attempted to represent null relative pointer",
		)
	}

	buf := bytes.NewBuffer(make([]byte, 0, 2))
	binary.Write(
		buf,
		binary.BigEndian,
		int16(relAddr.EntityPosition)-int16(relAddr.PointerPosition+1),
	)

	return buf.Bytes()
}

func (relAddr *relativeAddress) isNull() bool {
	return relAddr.PointerPosition == 0 || relAddr.EntityPosition == 0
}

type absoluteAddress struct {
	EntityPosition uint // where is pointed-to located
}

func newAbsAddr() *absoluteAddress {
	return &absoluteAddress{
		EntityPosition: 0,
	}
}

func (absAddr *absoluteAddress) evaluate(pos uint) {
	if absAddr.EntityPosition != 0 {
		logger.Printf(
			"warning: absolute pointer reevaluation (%d -> %d)",
			absAddr.EntityPosition,
			pos,
		)
	}

	absAddr.EntityPosition = pos
}

func (absAddr *absoluteAddress) represent() []byte {
	// null absolute pointer is OK, so no checking here

	buf := bytes.NewBuffer(make([]byte, 0, 2))
	binary.Write(
		buf,
		binary.BigEndian,
		int16(absAddr.EntityPosition),
	)

	return buf.Bytes()
}

func (absAddr *absoluteAddress) isNull() bool {
	return absAddr.EntityPosition == 0
}

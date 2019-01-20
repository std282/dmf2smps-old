package smpsbuild

import (
	"encoding/binary"
	"io"
)

type address interface {
	evaluate(pos uint)     // sets position of pointed-to
	represent(w io.Writer) // writes address in SMPS notation
	isNull() bool          // returns true if address doesn't point at anything
}

type relativeAddress struct {
	PointerPosition uint // where is pointer located
	EntityPosition  uint // where is pointed-to located
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

func (relAddr *relativeAddress) represent(w io.Writer) {
	if relAddr.isNull() {
		logger.Fatal(
			"error: attempted to represent null relative pointer",
		)
	}

	binary.Write(
		w,
		binary.BigEndian,
		int16(relAddr.EntityPosition)-int16(relAddr.PointerPosition+1),
	)
}

func (relAddr *relativeAddress) isNull() bool {
	return relAddr.PointerPosition == 0 || relAddr.EntityPosition == 0
}

type absoluteAddress struct {
	EntityPosition uint // where is pointed-to located
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

func (absAddr *absoluteAddress) represent(w io.Writer) {
	// null absolute pointer is OK, so no checking here

	binary.Write(
		w,
		binary.BigEndian,
		int16(absAddr.EntityPosition),
	)
}

func (absAddr *absoluteAddress) isNull() bool {
	return absAddr.EntityPosition == 0
}

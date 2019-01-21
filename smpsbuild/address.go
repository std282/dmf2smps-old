package smpsbuild

/*
	This file describes address interface and structures.

	Addresses are used for implementing pointers in SMPS, and they are lazily
	evaluated, since there is virtually no way to tell position of pointer and
	entity it points at until song structure is known.
*/

import (
	"encoding/binary"
	"io"
)

// address describes interface of address
type address interface {
	evaluate(pos uint)     // sets position of pointed-to
	represent(w io.Writer) // writes address in SMPS notation
}

// relativeAddress desribes address used in coordination flags. It points to
// some entity relatively to its own position.
type relativeAddress struct {
	pointerPosition uint // where is pointer located
	entityPosition  uint // where is pointed-to located
	active          bool // is pointer active
}

func (relAddr *relativeAddress) evaluate(pos uint) {
	if relAddr.active {
		logWarn.Printf(
			"relative pointer reevaluation (%d -> %d)",
			relAddr.entityPosition,
			pos,
		)
	}

	relAddr.entityPosition = pos
	relAddr.active = true
}

func (relAddr *relativeAddress) represent(w io.Writer) {
	if !relAddr.active {
		logError.Fatal(
			"attempted to represent inactive relative pointer",
		)
	}

	binary.Write(
		w,
		binary.BigEndian,
		int16(relAddr.entityPosition)-int16(relAddr.pointerPosition+1),
	)
}

// relativeAddress desribes address used in coordination flags. It points to
// some entity relatively to song start.
type absoluteAddress struct {
	entityPosition uint // where is pointed-to located
	active         bool
}

func (absAddr *absoluteAddress) evaluate(pos uint) {
	if absAddr.active {
		logWarn.Printf(
			"absolute pointer reevaluation (%d -> %d)",
			absAddr.entityPosition,
			pos,
		)
	}

	absAddr.entityPosition = pos
	absAddr.active = true
}

func (absAddr *absoluteAddress) represent(w io.Writer) {
	if !absAddr.active {
		logError.Fatal(
			"attempted to represent inactive absolute pointer",
		)
	}

	binary.Write(
		w,
		binary.BigEndian,
		int16(absAddr.entityPosition),
	)
}

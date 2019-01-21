package smpsbuild

/*
	This file describes DAC, FM and PSG headers and algorithms for their byte
	representation.
*/

import (
	"io"
)

// FM header
type fmHeader struct {
	dataPointer  *absoluteAddress
	boundPattern *Pattern
	volume       int8
	pitch        int8
}

func (fm *fmHeader) represent(w io.Writer) {
	fm.dataPointer.represent(w)
	w.Write([]byte{
		byte(fm.pitch),
		byte(fm.volume),
	})
}

func (*fmHeader) size() uint {
	return 4
}

// PSG header
type psgHeader struct {
	dataPointer  *absoluteAddress
	boundPattern *Pattern
	volume       int8
	pitch        int8
	initialVoice byte
}

func (psg *psgHeader) represent(w io.Writer) {
	psg.dataPointer.represent(w)
	w.Write([]byte{
		byte(psg.pitch),
		byte(psg.volume),
		byte(psg.initialVoice),
	})
}

func (*psgHeader) size() uint {
	return 5
}

// DAC header
type dacHeader struct {
	dataPointer  *absoluteAddress
	boundPattern *Pattern
}

func (dac *dacHeader) represent(w io.Writer) {
	dac.dataPointer.represent(w)
	w.Write([]byte{0x00, 0x00})
}

func (*dacHeader) size() uint {
	return 4
}

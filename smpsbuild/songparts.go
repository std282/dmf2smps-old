package smpsbuild

import (
	"io"
)

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

func (fm *fmHeader) size() uint {
	return 4
}

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

func (psg *psgHeader) size() uint {
	return 5
}

type dacHeader struct {
	dataPointer  *absoluteAddress
	boundPattern *Pattern
}

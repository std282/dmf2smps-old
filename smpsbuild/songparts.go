package smpsbuild

import "bytes"

type fmHeader struct {
	dataPointer  *absoluteAddress
	boundPattern *Pattern
	volume       int8
	pitch        int8
}

func (fm *fmHeader) represent() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 4))

	buf.Write(fm.dataPointer.represent())
	buf.Write([]byte{
		byte(fm.pitch),
		byte(fm.volume),
	})

	return buf.Bytes()
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

func (psg *psgHeader) represent() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 5))
	buf.Write(psg.dataPointer.represent())
	buf.Write([]byte{
		byte(psg.pitch),
		byte(psg.volume),
		byte(psg.initialVoice),
	})

	return buf.Bytes()
}

func (psg *psgHeader) size() uint {
	return 5
}

type dacHeader struct {
	dataPointer  *absoluteAddress
	boundPattern *Pattern
}

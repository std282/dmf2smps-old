package smpsbuild

/*
	This file contains function for adding other SMPS coordination flags to
	pattern:

	$E6 - alter volume (FM)
	$EC - alter volume (PSG)
	$E9 - alter pitch
	$E1 - fine-tune
	$E0 - panning
*/

import "io"

// AlterVolumeFM adds $E6 (alter volume for FM) coordination flag.
func (pat *Pattern) AlterVolumeFM(alteration int) {
	if alteration < -127 || alteration > 127 {
		logError.Fatalf(
			"invalid FM volume alteration (%d)",
			alteration,
		)
	}

	fmVol := new(eventAlterVolFM)
	fmVol.amount = int8(alteration)

	pat.addEvent(fmVol)
}

type eventAlterVolFM struct {
	amount int8
}

func (volFm *eventAlterVolFM) represent(w io.Writer) {
	w.Write([]byte{
		0xE6,
		byte(volFm.amount),
	})
}

func (*eventAlterVolFM) size() uint {
	return 2
}

// AlterVolumePSG adds $EC (alter volume for PSG) coordination flag.
func (pat *Pattern) AlterVolumePSG(alteration int) {
	if alteration < -127 || alteration > 127 {
		logError.Fatalf(
			"invalid PSG volume alteration (%d)",
			alteration,
		)
	}

	psgVol := new(eventAlterVolPSG)
	psgVol.amount = int8(alteration)

	pat.addEvent(psgVol)
}

type eventAlterVolPSG struct {
	amount int8
}

func (volFm *eventAlterVolPSG) represent(w io.Writer) {
	w.Write([]byte{
		0xEC,
		byte(volFm.amount),
	})
}

func (*eventAlterVolPSG) size() uint {
	return 2
}

// AlterPitch adds $E9 (alter pitch) coordination flag.
func (pat *Pattern) AlterPitch(alteration int) {
	if alteration < -127 || alteration > 127 {
		logError.Fatalf(
			"invalid pitch alteration (%d)",
			alteration,
		)
	}

	pitch := new(eventAlterPitch)
	pitch.amount = int8(alteration)

	pat.addEvent(pitch)
}

type eventAlterPitch struct {
	amount int8
}

func (pitch *eventAlterPitch) represent(w io.Writer) {
	w.Write([]byte{
		0xE9,
		byte(pitch.amount),
	})
}

func (*eventAlterPitch) size() uint {
	return 2
}

// SetFineTune adds $E1 (fine-tune channel) coordination flag.
func (pat *Pattern) SetFineTune(displacement int) {
	if displacement < -127 || displacement > 127 {
		logError.Fatalf(
			"invalid fine-tune value (%d)",
			displacement,
		)
	}

	tune := new(eventFineTune)
	tune.value = int8(displacement)

	pat.addEvent(tune)
}

type eventFineTune struct {
	value int8
}

func (tune *eventFineTune) represent(w io.Writer) {
	w.Write([]byte{
		0xE1,
		byte(tune.value),
	})
}

func (*eventFineTune) size() uint {
	return 2
}

// SetPan adds $E0 (set panning) coordination flag.
//
// side argument can be one of the following: smpsbuild.PanLeft, smpsbuild.PanRight
// or smpsbuild.PanCenter.
func (pat *Pattern) SetPan(side byte) {
	pan := new(eventPan)
	pan.side = side

	pat.addEvent(pan)
}

type eventPan struct {
	side byte
}

const (
	// PanLeft = set panning to left
	PanLeft byte = 0x40
	// PanRight = set panning to right
	PanRight byte = 0x80
	// PanCenter = set panning to center
	PanCenter byte = 0xC0
)

func (pan *eventPan) represent(w io.Writer) {
	w.Write([]byte{
		0xE0,
		pan.side,
	})
}

func (*eventPan) size() uint {
	return 2
}

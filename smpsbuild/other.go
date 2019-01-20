package smpsbuild

import "io"

// AlterVolumeFM adds $E6 (alter volume for FM) coordination flag.
func (pat *Pattern) AlterVolumeFM(alteration int) {
	if alteration < -127 || alteration > 127 {
		logger.Fatalf(
			"error: invalid FM volume alteration (%d)",
			alteration,
		)
	}

	fmVol := new(eventAlterVolFM)
	fmVol.Amount = int8(alteration)

	pat.events.PushBack(fmVol)
}

type eventAlterVolFM struct {
	Amount int8
}

func (volFm *eventAlterVolFM) represent(w io.Writer) {
	w.Write([]byte{
		0xE6,
		byte(volFm.Amount),
	})
}

func (*eventAlterVolFM) size() uint {
	return 2
}

// AlterVolumePSG adds $EC (alter volume for PSG) coordination flag.
func (pat *Pattern) AlterVolumePSG(alteration int) {
	if alteration < -127 || alteration > 127 {
		logger.Fatalf(
			"error: invalid PSG volume alteration (%d)",
			alteration,
		)
	}

	psgVol := new(eventAlterVolPSG)
	psgVol.Amount = int8(alteration)

	pat.events.PushBack(psgVol)
}

type eventAlterVolPSG struct {
	Amount int8
}

func (volFm *eventAlterVolPSG) represent(w io.Writer) {
	w.Write([]byte{
		0xEC,
		byte(volFm.Amount),
	})
}

func (*eventAlterVolPSG) size() uint {
	return 2
}

// AlterPitch adds $E9 (alter pitch) coordination flag.
func (pat *Pattern) AlterPitch(alteration int) {
	if alteration < -127 || alteration > 127 {
		logger.Fatalf(
			"error: invalid pitch alteration (%d)",
			alteration,
		)
	}

	pitch := new(eventAlterPitch)
	pitch.Amount = int8(alteration)

	pat.events.PushBack(pitch)
}

type eventAlterPitch struct {
	Amount int8
}

func (pitch *eventAlterPitch) represent(w io.Writer) {
	w.Write([]byte{
		0xE9,
		byte(pitch.Amount),
	})
}

func (*eventAlterPitch) size() uint {
	return 2
}

// SetFineTune adds $E1 (fine-tune channel) coordination flag.
func (pat *Pattern) SetFineTune(displacement int) {
	if displacement < -127 || displacement > 127 {
		logger.Fatalf(
			"error: invalid fine-tune value (%d)",
			displacement,
		)
	}

	tune := new(eventFineTune)
	tune.Value = int8(displacement)

	pat.events.PushBack(tune)
}

type eventFineTune struct {
	Value int8
}

func (tune *eventFineTune) represent(w io.Writer) {
	w.Write([]byte{
		0xE1,
		byte(tune.Value),
	})
}

func (*eventFineTune) size() uint {
	return 2
}

// SetPan adds $E0 (set panning) coordination flag.
//
// side argument can be one of the following: smpsbuild.PanLeft, smpsbuild.PanRight
// or smpsbuild.PanCenter.
func (pat *Pattern) SetPan(side panSide) {
	pan := new(eventPan)
	pan.Side = side

	pat.events.PushBack(pan)
}

type eventPan struct {
	Side panSide
}

type panSide byte

const (
	// PanLeft = set panning to left
	PanLeft panSide = 0x40
	// PanRight = set panning to right
	PanRight panSide = 0x80
	// PanCenter = set panning to center
	PanCenter panSide = 0xC0
)

func (pan *eventPan) represent(w io.Writer) {
	w.Write([]byte{
		0xE0,
		byte(pan.Side),
	})
}

func (*eventPan) size() uint {
	return 2
}

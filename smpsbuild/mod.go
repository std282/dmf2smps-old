package smpsbuild

/*
	This file contains functions for adding coordination flags for modulation
	control:

	$F0 - setup modulation
	$F1 - resume modulation
	$F4 - stop modulation
*/

import "io"

// SetModulation adds $F0 (set modulation parameters) coordination flag.
func (pat *Pattern) SetModulation(delay, stretch, factor, magnitude int) {
	switch {
	case delay < 0 || delay > 255:
		logWarn.Printf(
			"invalid modulation delay value (%d); it will likely sound wrong in SMPS",
			delay,
		)

	case stretch < 0 || stretch > 255:
		logWarn.Printf(
			"invalid modulation stretch value (%d); it will likely sound wrong in SMPS",
			stretch,
		)

	case factor < -127 || factor > 127:
		logWarn.Printf(
			"invalid modulation factor value (%d); it will likely sound wrong in SMPS",
			factor,
		)
	case magnitude < -127 || magnitude > 127:
		logWarn.Printf(
			"invalid modulation magnitude value (%d); it will likely sound wrong in SMPS",
			magnitude,
		)
	}

	mod := new(eventMod)
	mod.delay = uint8(delay)
	mod.factor = int8(factor)
	mod.magnitude = int8(magnitude)
	mod.stretch = uint8(stretch)

	pat.addEvent(mod)
}

type eventMod struct {
	delay     uint8
	stretch   uint8
	factor    int8
	magnitude int8
}

func (mod *eventMod) represent(w io.Writer) {
	w.Write([]byte{
		0xF0,
		byte(mod.delay),
		byte(mod.stretch),
		byte(mod.factor),
		byte(mod.magnitude),
	})
}

func (mod *eventMod) size() uint {
	return 5
}

// ResumeModulation adds $F1 (continue last active modulation) coordination flag.
func (pat *Pattern) ResumeModulation() {
	modOn := new(eventModOn)

	pat.addEvent(modOn)
}

type eventModOn struct{}

func (mod *eventModOn) represent(w io.Writer) {
	w.Write([]byte{0xF1})
}

func (*eventModOn) size() uint {
	return 1
}

// StopModulation adds $F4 (stop modulation) coordination flag.
func (pat *Pattern) StopModulation() {
	modOff := new(eventModOff)

	pat.addEvent(modOff)
}

type eventModOff struct{}

func (mod *eventModOff) represent(w io.Writer) {
	w.Write([]byte{0xF4})
}

func (*eventModOff) size() uint {
	return 1
}

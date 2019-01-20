package smpsbuild

import "io"

// SetModulation adds $F0 (set modulation parameters) coordination flag.
func (pat *Pattern) SetModulation(delay, stretch, factor, magnitude int) {
	switch {
	case delay < 0 || delay > 255:
		logger.Fatalf(
			"error: invalid modulation delay value (%d)",
			delay,
		)

	case stretch < 0 || stretch > 255:
		logger.Fatalf(
			"error: invalid modulation stretch value (%d)",
			stretch,
		)

	case factor < -127 || factor > 127:
		logger.Fatalf(
			"error: invalid modulation factor value (%d)",
			factor,
		)
	case magnitude < -127 || magnitude > 127:
		logger.Fatalf(
			"error: invalid modulation magnitude value (%d)",
			magnitude,
		)
	}

	mod := new(eventMod)
	mod.Delay = uint8(delay)
	mod.Factor = int8(factor)
	mod.Magnitude = int8(magnitude)
	mod.Stretch = uint8(stretch)

	pat.events.PushBack(mod)
}

type eventMod struct {
	Delay     uint8
	Stretch   uint8
	Factor    int8
	Magnitude int8
}

func (mod *eventMod) represent(w io.Writer) {
	w.Write([]byte{
		0xF0,
		byte(mod.Delay),
		byte(mod.Stretch),
		byte(mod.Factor),
		byte(mod.Magnitude),
	})
}

func (mod *eventMod) size() uint {
	return 5
}

// ResumeModulation adds $F1 (continue last active modulation) coordination flag.
func (pat *Pattern) ResumeModulation() {
	modOn := new(eventModOn)

	pat.events.PushBack(modOn)
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

	pat.events.PushBack(modOff)
}

type eventModOff struct{}

func (mod *eventModOff) represent(w io.Writer) {
	w.Write([]byte{0xF4})
}

func (*eventModOff) size() uint {
	return 1
}

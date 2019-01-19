package smpsbuild

type eventMod struct {
	Delay     uint8
	Stretch   uint8
	Factor    int8
	Magnitude int8
}

func (mod *eventMod) represent() []byte {
	return []byte{
		0xF0,
		byte(mod.Delay),
		byte(mod.Stretch),
		byte(mod.Factor),
		byte(mod.Magnitude),
	}
}

func (mod *eventMod) size() uint {
	return 5
}

type eventModOn struct{}

func (mod *eventModOn) represent() []byte {
	return []byte{
		0xF1,
	}
}

func (*eventModOn) size() uint {
	return 1
}

type eventModOff struct{}

func (mod *eventModOff) represent() []byte {
	return []byte{
		0xF4,
	}
}

func (*eventModOff) size() uint {
	return 1
}

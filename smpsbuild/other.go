package smpsbuild

type eventAlterVolFM struct {
	Amount int8
}

func (volFm *eventAlterVolFM) represent() []byte {
	return []byte{
		0xE6,
		byte(volFm.Amount),
	}
}

func (*eventAlterVolFM) size() uint {
	return 2
}

type eventAlterVolPSG struct {
	Amount int8
}

func (volFm *eventAlterVolPSG) represent() []byte {
	return []byte{
		0xEC,
		byte(volFm.Amount),
	}
}

func (*eventAlterVolPSG) size() uint {
	return 2
}

type eventAlterPitch struct {
	Amount int8
}

func (pitch *eventAlterPitch) represent() []byte {
	return []byte{
		0xE9,
		byte(pitch.Amount),
	}
}

func (*eventAlterPitch) size() uint {
	return 2
}

type eventFineTune struct {
	Value int8
}

func (tune *eventFineTune) represent() []byte {
	return []byte{
		0xE1,
		byte(tune.Value),
	}
}

func (*eventFineTune) size() uint {
	return 2
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

func (pan *eventPan) represent() []byte {
	return []byte{
		0xE0,
		byte(pan.Side),
	}
}

func (*eventPan) size() uint {
	return 2
}

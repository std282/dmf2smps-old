package dmfns

const (
	// Classic effects

	// Arpeggio = arpeggio
	Arpeggio byte = 0x00
	// PortaUp = portamento up
	PortaUp byte = 0x01
	// PortaDown = portamento down
	PortaDown byte = 0x02
	// PortaNote = portamento to note
	PortaNote byte = 0x03
	// Vibrato = vibrato
	Vibrato byte = 0x04
	// VolSlidePortaNote = volume slide + continue portamento to note
	VolSlidePortaNote byte = 0x05
	// VolSlideVibrato = volume slide + continue vibrato
	VolSlideVibrato byte = 0x06
	// Tremolo = tremolo
	Tremolo byte = 0x07
	// Panning = panning
	Panning byte = 0x08
	// SetSpeed1 = set speed 1 (even)
	SetSpeed1 byte = 0x09
	// VolSlide = volume slide
	VolSlide byte = 0x0A
	// GoToPattern = jump to pattern
	GoToPattern byte = 0x0B
	// Retrigger = retrigger note
	Retrigger byte = 0x0C
	// BreakPattern = jump to next pattern
	BreakPattern byte = 0x0D
	// SetSpeed2 = set speed 2 (odd)
	SetSpeed2 byte = 0x0F

	// Extended effects

	// ArpSpeed = set arpeggio speed
	ArpSpeed byte = 0xE0
	// PortaUpFix = portamento up by amount of semitones
	PortaUpFix byte = 0xE1
	// PortaDownFix = portamento down by amount of semitones
	PortaDownFix byte = 0xE2
	// VibratoMode = set vibrato mode (0: normal, 1: up only, 2: down only)
	VibratoMode byte = 0xE3
	// VibratoDepth = set vibrato depth
	VibratoDepth byte = 0xE4
	// FineTune = set fine tune
	FineTune byte = 0xE5
	// SampleBank = set sample bank
	SampleBank byte = 0xEB
	// NoteCut = cut note after several frames
	NoteCut byte = 0xEC
	// NoteDelay = delay note after several frames
	NoteDelay byte = 0xED

	// Chip-specific effects

	// EnableDAC = enable DAC on FM6
	EnableDAC byte = 0x17
	// NoiseMode = set noise mode on STD4
	NoiseMode byte = 0x20
)

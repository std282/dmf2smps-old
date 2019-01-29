package dmfns

const (
	// Classic effects

	// Arpeggio = arpeggio
	Arpeggio int16 = 0x00
	// PortaUp = portamento up
	PortaUp int16 = 0x01
	// PortaDown = portamento down
	PortaDown int16 = 0x02
	// PortaNote = portamento to note
	PortaNote int16 = 0x03
	// Vibrato = vibrato
	Vibrato int16 = 0x04
	// VolSlidePortaNote = volume slide + continue portamento to note
	VolSlidePortaNote int16 = 0x05
	// VolSlideVibrato = volume slide + continue vibrato
	VolSlideVibrato int16 = 0x06
	// Tremolo = tremolo
	Tremolo int16 = 0x07
	// Panning = panning
	Panning int16 = 0x08
	// SetSpeed1 = set speed 1 (even)
	SetSpeed1 int16 = 0x09
	// VolSlide = volume slide
	VolSlide int16 = 0x0A
	// GoToPattern = jump to pattern
	GoToPattern int16 = 0x0B
	// Retrigger = retrigger note
	Retrigger int16 = 0x0C
	// BreakPattern = jump to next pattern
	BreakPattern int16 = 0x0D
	// SetSpeed2 = set speed 2 (odd)
	SetSpeed2 int16 = 0x0F

	// Extended effects

	// ArpSpeed = set arpeggio speed
	ArpSpeed int16 = 0xE0
	// PortaUpFix = portamento up by amount of semitones
	PortaUpFix int16 = 0xE1
	// PortaDownFix = portamento down by amount of semitones
	PortaDownFix int16 = 0xE2
	// VibratoMode = set vibrato mode (0: normal, 1: up only, 2: down only)
	VibratoMode int16 = 0xE3
	// VibratoDepth = set vibrato depth
	VibratoDepth int16 = 0xE4
	// FineTune = set fine tune
	FineTune int16 = 0xE5
	// SampleBank = set sample bank
	SampleBank int16 = 0xEB
	// NoteCut = cut note after several frames
	NoteCut int16 = 0xEC
	// NoteDelay = delay note after several frames
	NoteDelay int16 = 0xED

	// Chip-specific effects

	// EnableDAC = enable DAC on FM6
	EnableDAC int16 = 0x17
	// NoiseMode = set noise mode on STD4
	NoiseMode int16 = 0x20
)

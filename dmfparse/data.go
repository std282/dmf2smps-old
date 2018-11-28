package dmfparse

// Song represents DefleMask Format (DMF) module
type Song struct {
	Name            string
	Author          string
	TimeBase        int
	TickTime1       int
	TickTime2       int
	FramesPerSecond int

	Matrix      [][]int
	Instruments []Instrument
	Channels    [10][]Pattern
	Samples     []Sample
}

// Instrument describes common DMF instrument types
type Instrument interface {
	Name() string
	Type() InstType
}

// InstType represents type of instrument, either FM or STD
type InstType byte

const (
	// FM = FM instrument type
	FM InstType = 1
	// STD = STD instrument type
	STD InstType = 0
)

// InstrumentFM represents FM instrument
type InstrumentFM struct {
	name string

	ALG, FB, LFO, LFO2 byte

	AM, AR, DR, MULT, RR, SL, TL, DT2, RS, D2R, SSG [4]byte
}

// Name returns FM instrument name
func (fm *InstrumentFM) Name() string {
	return fm.name
}

// Type returns InstType constant
func (*InstrumentFM) Type() InstType {
	return FM
}

// InstrumentSTD represents STD instrument
type InstrumentSTD struct {
	name string

	VolumeEnv, ArpeggioEnv, NoiseEnv    []int32
	VolumeLoop, ArpeggioLoop, NoiseLoop bool
	ArpeggioMode                        bool
}

// Name returns STD instrument name
func (std *InstrumentSTD) Name() string {
	return std.name
}

// Type returns InstType constant
func (*InstrumentSTD) Type() InstType {
	return STD
}

// Pattern represents DMF pattern
type Pattern struct {
	effectsAmount int
	Rows          []Row
}

// Row represents DMF row
type Row struct {
	parent *Pattern

	Note    int16
	Octave  int16
	Volume  int16
	InstNum int16
	Effects []Effect
}

// EffectsAmount returns amount of effects in the row
func (row *Row) EffectsAmount() int {
	return row.parent.effectsAmount
}

// Effect represents single DMF effect
type Effect struct {
	Type int16
	Byte int16
}

// Sample represents a DMF sample
type Sample struct {
	Name string
}

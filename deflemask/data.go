package deflemask

// Module represents single DefleMask song.
type Module struct {
	Name   string
	Author string

	BaseTime  int
	TimeTick1 int
	TimeTick2 int
	FPS       int

	PatternCount    int
	PatternLength   int
	InstrumentCount int
	SampleCount     int

	patternMatrix [10][]int
	songData      [10][][]Row
	instruments   []interface{}
	sampleNames   []string
}

// Row represents single DefleMask row.
type Row struct {
	Note       int16
	Volume     int16
	Instrument int16
	Effects    [4]Effect
}

// Effect represents single DefleMask effect.
type Effect struct {
	Type  int16
	Value int16
}

// InstrumentFM represents FM instrument in DefleMask
type InstrumentFM struct {
	Algorithm, Feedback int
	FreqMod, AmpMod     int

	AmpModFlag                              bool
	Attack, Decay, Sustain, Decay2, Release int
	Detune, Multiply, RateScale, TotalLevel int
	Detune2, SSGEnvelopeGenerator           int

	Name string
}

// InstrumentSTD represents STD instrument in DefleMask
type InstrumentSTD struct {
	Volume *Envelope
	Pitch  *Envelope
	Noise  *Envelope

	Name string
}

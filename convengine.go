package main

import "bytes"

// FMEngine is conversion engine for FM channels
type FMEngine struct {
	buf       bytes.Buffer // buffer for SMPS data
	state     FMState      // engine state
	acmFrames int          // accumulated frames
	rowLen    *int         // length of row in frames
}

// DACEngine is conversion engine for DAC channel
type DACEngine struct {
	buf       bytes.Buffer // buffer for SMPS data
	state     DACState     // engine state
	acmFrames int          // accumulated frames
	rowLen    *int         // length of row in frames
}

// PSGEngine is conversion engine for PSG tonal channel
type PSGEngine struct {
	buf       bytes.Buffer // buffer for SMPS data
	state     PSGState     // engine state
	acmFrames int          // accumulated frames
	rowLen    *int         // length of row in frames
}

// NoiseEngine is conversion engine for PSG noise channel
type NoiseEngine struct {
	buf       bytes.Buffer // buffer for SMPS data
	state     NoiseState   // engine state
	acmFrames int          // accumulated frames
	rowLen    *int         // length of row in frames
}

// FMState is the state of FM conversion engine
type FMState struct {
	lastNote   byte // last SMPS note played
	volume     int  // current SMPS volume
	instrument int  // current SMPS FM instrument
	freq       int  // current frequency
	freqBias   int  // frequency bias

	arp             arpState      // arpeggio state
	portaContinuous portaConState // continuous portamento state
	portaHoming     portaHomState // homing portamento state
	vibrato         LFO           // vibrato oscillator
	tremolo         LFO           // tremolo oscillator
	volSlide        volSlideState // volume slide state

	vibratoMode byte // vibrato mode
}

// DACState is the state of DAC conversion engine
type DACState struct {
	lastSample byte // last DAC sample played
}

// PSGState is the state of PSG conversion engine
type PSGState struct {
	lastNote   byte // last SMPS note played
	volume     int  // SMPS volume
	instrument int  // current SMPS PSG volume envelope
	freq       int  // current frequency
	freqBias   int  // frequency bias

	arp             arpState      // arpeggio state
	portaContinuous portaConState // continuous portamento state
	portaHoming     portaHomState // homing portamento state
	vibrato         LFO           // vibrato oscillator
	tremolo         LFO           // tremolo oscillator
	volSlide        volSlideState // volume slide state

	vibratoMode byte // vibrato mode

	instVECount int // volume envelope counter
	instPECount int // pitch envelope counter
}

// NoiseState is the state of PSG noise conversion engine
type NoiseState struct {
	lastNote   byte // last SMPS note played
	volume     int  // SMPS volume
	instrument int  // current SMPS PSG volume envelope
	freq       int  // current frequency
	freqBias   int  // frequency bias
	noiseMode  byte // noise mode

	arp             arpState      // arpeggio state
	portaContinuous portaConState // continuous portamento state
	portaHoming     portaHomState // homing portamento state
	vibrato         LFO           // vibrato oscillator
	tremolo         LFO           // tremolo oscillator
	volSlide        volSlideState // volume slide state

	vibratoMode byte // vibrato mode

	instVECount int // volume envelope counter
	instPECount int // pitch envelope counter
}

// Arpeggio (00xy) state
type arpState struct {
	disp1   int
	disp2   int
	cycleNo int
	length  int
	frameNo int
}

// Continuous portamento (01xx, 02xx) state
type portaConState struct {
	magnitude int
	frameNo   int
}

// Homing portamento (03xx, E1xy, E2xy) state
type portaHomState struct {
	magnitude int
	remained  int
}

// Volume slide (0Axy) state
type volSlideState struct {
	magnitude int
	residue   int
}

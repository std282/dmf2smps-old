package smpsbuild

/*
	This file contains functions for adding coordination flags for voice setting:

	$EF - set FM voice
	$F5 - set PSG voice (volume envelope)
	$F3 - set PSG3 noise mode
*/

import "io"

// SetFMVoice adds $EF (set FM voice) coordination flag.
func (pat *Pattern) SetFMVoice(voiceNum int) {
	if voiceNum < 0 || voiceNum > 255 {
		logError.Fatalf(
			"invalid FM voice number (%d);",
			voiceNum,
		)
	}

	fmVoice := new(eventSetVoiceFM)
	fmVoice.voicePos = uint8(voiceNum)

	pat.addEvent(fmVoice)
}

type eventSetVoiceFM struct {
	voicePos uint8
}

func (fmVoice *eventSetVoiceFM) represent(w io.Writer) {
	w.Write([]byte{
		0xEF,
		byte(fmVoice.voicePos),
	})
}

func (*eventSetVoiceFM) size() uint {
	return 2
}

// SetPSGVoice adds $F5 (set PSG voice) coordination flag.
func (pat *Pattern) SetPSGVoice(voiceNum int) {
	if voiceNum < 0 || voiceNum > 255 {
		logError.Fatalf(
			"error: invalid PSG voice number (%d)",
			voiceNum,
		)
	}

	psgVoice := new(eventSetVoicePSG)
	psgVoice.voicePos = uint8(voiceNum)

	pat.addEvent(psgVoice)
}

type eventSetVoicePSG struct {
	voicePos uint8
}

func (psgVoice *eventSetVoicePSG) represent(w io.Writer) {
	w.Write([]byte{
		0xF5,
		byte(psgVoice.voicePos),
	})
}

func (*eventSetVoicePSG) size() uint {
	return 2
}

// SetNoiseMode adds $F3 (set PSG noise mode) coordination flag.
//
// Use values WaveRandom or WavePeriodic for wave.
//
// Use values RangeLowOnly, RangeMidOnly, RangeHighOnly or RangeFull for rng.
func (pat *Pattern) SetNoiseMode(wave noiseWave, rng noiseRange) {
	noise := new(eventPsgNoise)
	noise.nRange = rng
	noise.wave = wave

	pat.addEvent(noise)
}

type eventPsgNoise struct {
	wave   noiseWave
	nRange noiseRange
}

type noiseWave byte
type noiseRange byte

const (
	// RangeLowOnly = single low note for noise
	RangeLowOnly noiseRange = iota
	// RangeMidOnly = single middle note for noise
	RangeMidOnly
	// RangeHighOnly = single high note for noise
	RangeHighOnly
	// RangeFull = full range of notes for noise
	RangeFull

	// WavePeriodic = periodic noise (low-duty pulse wave)
	WavePeriodic noiseWave = 0xE0
	// WaveRandom = random noise
	WaveRandom noiseWave = 0xE4
)

func (noise *eventPsgNoise) represent(w io.Writer) {
	w.Write([]byte{
		0xF3,
		byte(noise.wave) | byte(noise.nRange),
	})
}

func (*eventPsgNoise) size() uint {
	return 2
}

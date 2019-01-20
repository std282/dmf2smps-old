package smpsbuild

import "io"

// SetFMVoice adds $EF (set FM voice) coordination flag.
func (pat *Pattern) SetFMVoice(voiceNum int) {
	if voiceNum < 0 || voiceNum > 255 {
		logger.Printf(
			"error: invalid FM voice number (%d)",
			voiceNum,
		)
	}

	fmVoice := new(eventSetVoiceFM)
	fmVoice.VoicePos = uint8(voiceNum)

	pat.events.PushBack(fmVoice)
}

type eventSetVoiceFM struct {
	VoicePos uint8
}

func (fmVoice *eventSetVoiceFM) represent(w io.Writer) {
	w.Write([]byte{
		0xEF,
		byte(fmVoice.VoicePos),
	})
}

func (*eventSetVoiceFM) size() uint {
	return 2
}

// SetPSGVoice adds $F5 (set PSG voice) coordination flag.
func (pat *Pattern) SetPSGVoice(voiceNum int) {
	if voiceNum < 0 || voiceNum > 255 {
		logger.Printf(
			"error: invalid PSG voice number (%d)",
			voiceNum,
		)
	}

	psgVoice := new(eventSetVoicePSG)
	psgVoice.VoicePos = uint8(voiceNum)

	pat.events.PushBack(psgVoice)
}

type eventSetVoicePSG struct {
	VoicePos uint8
}

func (psgVoice *eventSetVoicePSG) represent(w io.Writer) {
	w.Write([]byte{
		0xF5,
		byte(psgVoice.VoicePos),
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
	noise.Range = rng
	noise.Wave = wave

	pat.events.PushBack(noise)
}

type eventPsgNoise struct {
	Wave  noiseWave
	Range noiseRange
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
		byte(noise.Wave) | byte(noise.Range),
	})
}

package smpsbuild

type eventSetVoiceFM struct {
	VoicePos uint8
}

func (fmVoice *eventSetVoiceFM) represent() []byte {
	return []byte{
		0xEF,
		byte(fmVoice.VoicePos),
	}
}

func (*eventSetVoiceFM) size() uint {
	return 2
}

type eventSetVoicePSG struct {
	VoicePos uint8
}

func (psgVoice *eventSetVoicePSG) represent() []byte {
	return []byte{
		0xF5,
		byte(psgVoice.VoicePos),
	}
}

func (*eventSetVoicePSG) size() uint {
	return 2
}

type eventPsgNoise struct {
	Wave  noiseWave
	Range noiseRange
}

type noiseWave byte
type noiseRange byte

const (
	// NoiseLowOnly = single low note for noise
	NoiseLowOnly noiseRange = iota
	// NoiseMidOnly = single middle note for noise
	NoiseMidOnly
	// NoiseHighOnly = single high note for noise
	NoiseHighOnly
	// NoiseFullRange = full range of notes for noise
	NoiseFullRange

	// ModePeriodic = periodic noise (low-duty pulse wave)
	ModePeriodic noiseWave = 0xE0
	// ModeRandom = random noise
	ModeRandom noiseWave = 0xE4
)

func (noise *eventPsgNoise) represent() []byte {
	return []byte{
		0xF3,
		byte(noise.Wave) | byte(noise.Range),
	}
}

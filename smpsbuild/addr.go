package smpsbuild

import "errors"

// SetupChannels sets song up for filling
func (song *Song) SetupChannels(fm int, psg int) {
	song.channelsFM = fm
	song.channelsPSG = psg

	song.offsetFM = make([]AbsAddress, fm-1)
	song.pitchFM = make([]int, fm-1)
	song.volumeFM = make([]int, fm-1)

	song.offsetPSG = make([]AbsAddress, psg)
	song.pitchPSG = make([]int, psg)
	song.volumePSG = make([]int, psg)
	song.voicePSG = make([]int, psg)
}

// SetupTempo sets up SMPS tempo
func (song *Song) SetupTempo(div int, mod int) {
	song.tempoDivider = div
	song.tempoModifier = mod
}

// AddVoice adds one more voice to the song
func (song *Song) AddVoice(vc Voice) {
	song.voices = append(song.voices, vc)
}

// SetFMInitParams sets FM channel initial parameters
func (song *Song) SetFMInitParams(fm channel, vol int, pitch int) {
	pos := int(fm) - int(FM1)
	if pos < 0 || pos > 5 {
		panic(errors.New("addressed non-FM channel when expected FM"))
	}

	song.volumeFM[pos] = vol
	song.pitchFM[pos] = pitch
}

// SetPSGInitParams sets PSG channel initial parameters
func (song *Song) SetPSGInitParams(psg channel, vol int, pitch int, env int) {
	pos := int(psg) - int(PSG1)
	if pos < 0 {
		panic(errors.New("addressed non-PSG channel when expected PSG"))
	}

	song.volumePSG[pos] = vol
	song.pitchPSG[pos] = pitch
	song.voicePSG[pos] = env
}

package smpsbuild

import (
	"container/list"
	"io"
)

// Song represents SMPS song
type Song struct {
	// Main header data
	voices        []Voice // FM voices
	tempoDivider  int     // self-explanatory
	tempoModifier int     // self-explanatory
	channelsFM    int     // amount of FM channels + DAC
	channelsPSG   int     // amount of PSG channels

	// Channel header data
	offsetDAC absAddress   // pointer to initial DAC pattern
	offsetFM  []absAddress // pointers to initial FM patterns
	volumeFM  []int        // initial volume of FM channels
	pitchFM   []int        // initial pitch of FM channels
	offsetPSG []absAddress // pointers to initial PSG patterns
	volumePSG []int        // initial volume of PSG channels
	pitchPSG  []int        // initial pitch of PSG channels
	voicePSG  []int        // initial voice of PSG channels

	// Note data
	data list.List // patterns list
}

// NewSong creates a new valid SMPS song
func NewSong() (song *Song) {
	return &Song{tempoDivider: 1}
}

// Voice represents SMPS FM voice
type Voice struct {
	/* Feedback, Algorithm */
	FB, ALG int

	/* For operators 1, 3, 2, 4:
	 * -- Frequency multiplier
	 * -- Detune amount
	 * -- Attack rate
	 * -- Rate scaling factor
	 * -- Detune rate
	 * -- Sustain rate
	 * -- Release rate
	 * -- Sustain level
	 * -- Total level
	 */
	MULT, DT, AR, RS, DR, SR, RR, SL, TL [4]int
}

// Assemble exports SMPS song to binary
func (song *Song) Assemble(w io.Writer) {
	song.resolveAddresses()
	song.export(w)
}

// SetupChannels sets song up for filling
func (song *Song) SetupChannels(fm int, psg int) {
	song.channelsFM = fm
	song.channelsPSG = psg

	// DAC is considered FM channel; so there is one less FM channel than declared
	song.offsetFM = make([]absAddress, fm-1)
	song.pitchFM = make([]int, fm-1)
	song.volumeFM = make([]int, fm-1)

	song.offsetPSG = make([]absAddress, psg)
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
	pos := int(fm - FM1)
	if pos < 0 || pos > 5 {
		logger.Fatal("error: addressed non-FM channel when expected FM")
	}

	song.volumeFM[pos] = vol
	song.pitchFM[pos] = pitch
}

// SetPSGInitParams sets PSG channel initial parameters
func (song *Song) SetPSGInitParams(psg channel, vol int, pitch int, env int) {
	pos := int(psg - PSG1)
	if pos < 0 {
		logger.Fatal("error: addressed non-PSG channel when expected PSG")
	}

	song.volumePSG[pos] = vol
	song.pitchPSG[pos] = pitch
	song.voicePSG[pos] = env
}

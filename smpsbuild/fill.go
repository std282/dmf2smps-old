package smpsbuild

// SetupChannels sets song up for filling
func (song *Song) SetupChannels(fm int, psg int) {
	song.channelsFM = fm
	song.channelsPSG = psg

	song.offsetFM = make([]AbsAddress, fm)
	song.pitchFM = make([]int, fm)
	song.volumeFM = make([]int, fm)

	song.offsetPSG = make([]AbsAddress, psg)
	song.pitchPSG = make([]int, psg)
	song.volumePSG = make([]int, psg)
	song.voicePSG = make([]int, psg)
}

// AddVoice adds one more voice to the song
func (song *Song) AddVoice(vc *Voice) {
	song.voices = append(song.voices, *vc)
}

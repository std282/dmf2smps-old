package smpsbuild

// channel represents a channel ID
type channel int

const (
	// DAC channel
	DAC channel = iota
	// FM1 channel
	FM1
	// FM2 channel
	FM2
	// FM3 channel
	FM3
	// FM4 channel
	FM4
	// FM5 channel
	FM5
	// FM6 channel
	FM6
	// PSG1 channel
	PSG1
	// PSG2 channel
	PSG2
	// PSG3 channel
	PSG3
)

// InitChannel sets pattern that channel starts to play from
func (song *Song) InitChannel(c channel, pat *Pattern) {
	switch c {
	case DAC:
		song.offsetDAC.refer(pat)

	case FM1, FM2, FM3, FM4, FM5, FM6:
		song.offsetFM[int(c-FM1)].refer(pat)

	case PSG1, PSG2, PSG3:
		song.offsetPSG[int(c-PSG1)].refer(pat)
	}
}

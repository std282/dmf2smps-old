package smpsbuild

import (
	"container/list"
)

// Song represents SMPS song
type Song struct {
	voicePtr   *absoluteAddress
	tempoDiv   uint8
	tempoMod   uint8
	dacHeader  dacHeader
	fmHeaders  [6]fmHeader
	psgHeaders [3]psgHeader

	voices   *list.List
	noteData *list.List
}

// NewSong returns new SMPS Song ready to use
func NewSong() *Song {
	newSongPtr := &Song{
		voicePtr: newAbsAddr(),
		tempoDiv: 0,
		tempoMod: 0,
		voices:   list.New(),
		noteData: list.New(),
	}

	newSongPtr.dacHeader.dataPointer = newAbsAddr()

	for i := range newSongPtr.fmHeaders {
		newSongPtr.fmHeaders[i].dataPointer = newAbsAddr()
	}

	for i := range newSongPtr.psgHeaders {
		newSongPtr.psgHeaders[i].dataPointer = newAbsAddr()
	}

	return newSongPtr
}

// SetTempo sets SMPS song tempo to specified values
func (song *Song) SetTempo(div, mod int) {
	song.tempoDiv = uint8(div)
	song.tempoMod = uint8(mod)
}

// SetInitialPattern sets initial pattern for channel and adds it to data list
func (song *Song) SetInitialPattern(pat *Pattern, ch channelID) {
	setRef := func(workPat *Pattern, addr *absoluteAddress) {
		if workPat != nil {
			workPat.removeRef(addr)
		}
		workPat = pat
		song.noteData.PushBack(pat)
		pat.addRef(addr)
	}

	switch ch {
	case DAC:
		setRef(
			song.dacHeader.boundPattern,
			song.dacHeader.dataPointer,
		)

	case FM1:
		setRef(
			song.fmHeaders[0].boundPattern,
			song.fmHeaders[0].dataPointer,
		)

	case FM2:
		setRef(
			song.fmHeaders[1].boundPattern,
			song.fmHeaders[1].dataPointer,
		)

	case FM3:
		setRef(
			song.fmHeaders[2].boundPattern,
			song.fmHeaders[2].dataPointer,
		)

	case FM4:
		setRef(
			song.fmHeaders[3].boundPattern,
			song.fmHeaders[3].dataPointer,
		)

	case FM5:
		setRef(
			song.fmHeaders[4].boundPattern,
			song.fmHeaders[4].dataPointer,
		)

	case FM6:
		setRef(
			song.fmHeaders[5].boundPattern,
			song.fmHeaders[5].dataPointer,
		)

	case PSG1:
		setRef(
			song.psgHeaders[0].boundPattern,
			song.psgHeaders[0].dataPointer,
		)

	case PSG2:
		setRef(
			song.psgHeaders[1].boundPattern,
			song.psgHeaders[1].dataPointer,
		)

	case PSG3:
		setRef(
			song.psgHeaders[2].boundPattern,
			song.psgHeaders[2].dataPointer,
		)
	}
}

type channelID int

const (
	// DAC channel
	DAC channelID = iota
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

// AttachPattern attaches single pattern to song
func (song *Song) AttachPattern(pat *Pattern) {
	song.noteData.PushBack(pat)
}

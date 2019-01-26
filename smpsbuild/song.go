package smpsbuild

/*
	This file describes SMPS song structure and its interface.
*/

import (
	"container/list"
	"io"
)

// Song represents SMPS song
type Song struct {
	voicePtr   *absoluteAddress
	fmAmount   uint8
	psgAmount  uint8
	tempoDiv   uint8
	tempoMod   uint8
	dacHeader  dacHeader
	fmHeaders  [6]fmHeader
	psgHeaders [3]psgHeader

	voices      *list.List
	noteData    *list.List
	stopPattern *Pattern
}

// NewSong returns new SMPS song ready to use
func NewSong() *Song {
	song := &Song{
		voicePtr: new(absoluteAddress),
		voices:   list.New(),
		noteData: list.New(),
	}

	// Following instructions initialize absolute addresses for SMPS headers.

	song.dacHeader.dataPointer = new(absoluteAddress)

	for i := range song.fmHeaders {
		song.fmHeaders[i].dataPointer = new(absoluteAddress)
	}

	for i := range song.psgHeaders {
		song.psgHeaders[i].dataPointer = new(absoluteAddress)
	}

	// Each song has stop pattern. It is used for disabling channels if there
	// are channels that do not have any data to play, but there are channels
	// with data after them.
	//
	// For example, PSG3 noise channel must play, even if there are no notes for
	// PSG1 and PSG2.

	song.stopPattern = NewPattern()
	song.stopPattern.PlaceStop()

	// set all patterns to stop
	for ch := DAC; ch <= PSG3; ch++ {
		_, addr := song.setInitGetAddr(ch, song.stopPattern)
		song.stopPattern.addRef(addr)
	}

	return song
}

// SetTempo sets SMPS song tempo to specified values
func (song *Song) SetTempo(div, mod int) {
	song.tempoDiv = uint8(div)
	song.tempoMod = uint8(mod)
}

// SetInitialPattern sets initial pattern for channel and attaches it to pattern
// list.
func (song *Song) SetInitialPattern(pat *Pattern, ch channelID) {
	oldPat, addr := song.setInitGetAddr(ch, pat)
	oldPat.removeRef(addr)
	pat.addRef(addr)
	song.noteData.PushBack(pat)
}

// setInitGetAddr sets boundPattern value of specified channel to specified
// pattern. It returns pointer to old pattern that was bound to channel.
//
// Please note: this function does NOT add new pattern to list. It is designed
// to work with patterns that are already in the list.
func (song *Song) setInitGetAddr(ch channelID, pat *Pattern) (*Pattern, *absoluteAddress) {
	switch ch {
	case DAC:
		initPattern := song.dacHeader.boundPattern
		initAddress := song.dacHeader.dataPointer

		song.dacHeader.boundPattern = pat
		return initPattern, initAddress

	case FM1, FM2, FM3, FM4, FM5, FM6:
		index := int(ch - FM1)
		initPattern := song.fmHeaders[index].boundPattern
		initAddress := song.fmHeaders[index].dataPointer

		song.fmHeaders[index].boundPattern = pat
		return initPattern, initAddress

	case PSG1, PSG2, PSG3:
		index := int(ch - PSG1)
		initPattern := song.psgHeaders[index].boundPattern
		initAddress := song.psgHeaders[index].dataPointer

		song.psgHeaders[index].boundPattern = pat
		return initPattern, initAddress

	// this should never happen
	default:
		return nil, nil
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

// AttachPattern attaches single pattern to song.
//
// Don't attach same pattern twice. If you do, you'll likely get warnings about
// pointer reevaluation.
func (song *Song) AttachPattern(pat *Pattern) {
	song.noteData.PushBack(pat)
}

// AttachVoice attaches single FM voice to song.
func (song *Song) AttachVoice(voice *Voice) {
	song.voices.PushBack(voice)
}

// Export exports SMPS song to writer
func (song *Song) Export(w io.Writer) (err error) {
	song.determineChannels()
	song.removeUnusedPatterns()
	song.settleReferences()

	pw := newPanicWriter(w)
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	song.export(pw)
	return
}

package mngpat

import (
	"github.com/std282/dmf2smps/smpsbuild"
)

// PlaceNote adds note to pattern with SMPS optimization
func (mpat *ManagedPattern) PlaceNote(note byte, length int) {
	isSameNote := (note == smpsbuild.SameNote) || note == mpat.note
	isSameLength := (length == smpsbuild.SameLength) || length == int(mpat.noteLength)

	var actualNote byte
	if isSameNote {
		actualNote = mpat.note
	} else {
		actualNote = note
	}

	var actualLength int
	if isSameLength {
		actualLength = int(mpat.noteLength)
	} else {
		actualLength = length
	}

	if isSameNote {
		mpat.placeNoteEx(smpsbuild.SameNote, actualLength)
	} else if isSameLength {
		mpat.placeNoteEx(actualNote, smpsbuild.SameLength)
	} else {
		mpat.placeNoteEx(actualNote, actualLength)
	}
}

func (mpat *ManagedPattern) placeNoteEx(note byte, length int) {
	var lastLength int8
	if length > 127 {
		mpat.innerPattern.PlaceNote(note, 127)
		length -= 127
		for length > 127 {
			mpat.innerPattern.PreventAttack()
			mpat.innerPattern.PlaceNote(smpsbuild.SameNote, 127)
			length -= 127
		}

		lastLength = 127

		if length > 0 {
			mpat.innerPattern.PreventAttack()
			mpat.innerPattern.PlaceNote(smpsbuild.SameNote, length)
			lastLength = int8(length)
		}
	} else {
		mpat.innerPattern.PlaceNote(note, length)
		lastLength = int8(length)
	}

	if note != smpsbuild.SameNote {
		mpat.note = note
	}

	if length != smpsbuild.SameLength {
		mpat.noteLength = lastLength
	}
}

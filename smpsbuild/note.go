package smpsbuild

/*
	This file contains function for adding notes to pattern.
*/

import "io"

// PlaceNote places SMPS note of specified length to pattern.
//
// Use smpsbuild.SameNote (or 0xFF) as note argument to play last played note.
//
// Use smpsbuild.SameLength (or -1) as length argument to play note with last
// length.
func (pat *Pattern) PlaceNote(note byte, length int) {
	if note == SameNote && length == SameLength {
		logError.Fatal(
			"invalid SMPS note: (SameNote, SameLength)",
		)
	}

	if (note < 0x80 || note > 0xDF) && note != SameNote {
		logError.Fatalf(
			"invalid SMPS note value (%d)",
			note,
		)
	}

	if (length <= 0 && length != SameLength) || length > 127 {
		logError.Fatalf(
			"invalid SMPS note length (%d)",
			length,
		)
	}

	noteEv := new(eventNote)
	noteEv.note = note
	noteEv.length = int8(length)

	pat.addEvent(noteEv)
}

// Note event
type eventNote struct {
	note   byte
	length int8
}

// SameNote indicates that note to be played is the same as previous
const SameNote byte = 0xFF

// SameLength indicated that note to be played will be of same length
const SameLength int = -1
const sameLengthInternal int8 = -1

// Returns byte representation of note
func (note *eventNote) represent(w io.Writer) {
	switch {
	case note.note == SameNote:
		w.Write([]byte{
			byte(note.length),
		})

	case note.length == sameLengthInternal:
		w.Write([]byte{
			note.note,
		})

	default:
		w.Write([]byte{
			note.note,
			byte(note.length),
		})
	}
}

// Returns size of note
func (note *eventNote) size() uint {
	if note.note == SameNote || note.length == sameLengthInternal {
		return 1
	}

	return 2
}

// PreventAttack adds $E7 (prevent next note from attacking) coordination flag.
func (pat *Pattern) PreventAttack() {
	noAtk := new(eventNoAttack)

	pat.addEvent(noAtk)
}

type eventNoAttack struct{}

func (*eventNoAttack) represent(w io.Writer) {
	w.Write([]byte{0xE7})
}

func (*eventNoAttack) size() uint {
	return 1
}

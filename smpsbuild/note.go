package smpsbuild

import "io"

// PlaceNote places SMPS note of specified length to pattern.
//
// Use smpsbuild.SameNote (or 0xFF) as note argument to play last played note.
//
// Use smpsbuild.SameLength (or -1) as length argument to play note with last
// length.
func (pat *Pattern) PlaceNote(note byte, length int) {
	if (note < 0x80 || note > 0xDF) && note != SameNote {
		logger.Fatalf(
			"error: invalid SMPS note value (%d)",
			note,
		)
	}

	if (length <= 0 && length != SameLength) || length > 127 {
		logger.Fatalf(
			"error: invalid SMPS note length (%d)",
			length,
		)
	}

	noteEv := new(eventNote)
	noteEv.Note = note
	noteEv.Length = int8(length)

	if note != SameNote {
		pat.lastNote = noteEv.Note
	}

	if length != SameLength {
		pat.lastLength = noteEv.Length
	}

	pat.events.PushBack(noteEv)
}

// Note event
type eventNote struct {
	Note   byte
	Length int8
}

// SameNote indicates that note to be played is the same as previous
const SameNote byte = 0xFF

// SameLength indicated that note to be played will be of same length
const SameLength int = -1
const sameLengthInternal int8 = -1

// Returns byte representation of note
func (note *eventNote) represent(w io.Writer) {
	switch {
	case note.Note == SameNote:
		w.Write([]byte{
			byte(note.Length),
		})

	case note.Length == sameLengthInternal:
		w.Write([]byte{
			note.Note,
		})

	default:
		w.Write([]byte{
			note.Note,
			byte(note.Length),
		})
	}
}

// Returns size of note
func (note *eventNote) size() uint {
	if note.Note == SameNote || note.Length == sameLengthInternal {
		return 1
	}

	return 2
}

// PlaceNoteClever places note with awareness of last note and last length.
// If note is the same as last, places only length.
// If length is the same as last, places only note.
func (pat *Pattern) PlaceNoteClever(note byte, length int) {
	switch {
	case note == pat.lastNote:
		pat.PlaceNote(SameNote, length)

	case length == int(pat.lastLength):
		pat.PlaceNote(note, SameLength)

	default:
		pat.PlaceNote(note, length)
	}
}

// PreventAttack adds $E7 (prevent next note from attacking) coordination flag.
func (pat *Pattern) PreventAttack() {
	noAtk := new(eventNoAttack)

	pat.events.PushBack(noAtk)
}

type eventNoAttack struct{}

func (*eventNoAttack) represent(w io.Writer) {
	w.Write([]byte{0xE7})
}

func (*eventNoAttack) size() uint {
	return 1
}

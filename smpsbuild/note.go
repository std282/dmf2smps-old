package smpsbuild

// Note event
type eventNote struct {
	Note   byte
	Length int8
}

// SameNote indicates that note to be played is the same as previous
const SameNote byte = 0xFF

// SameLength indicated that note to be played will be of same length
const SameLength int8 = -1

// Returns byte representation of note
func (note *eventNote) represent() []byte {
	switch {
	case note.Note == SameNote:
		return []byte{
			byte(note.Length),
		}

	case note.Length == SameLength:
		return []byte{
			note.Note,
		}

	default:
		return []byte{
			note.Note,
			byte(note.Length),
		}
	}
}

// Returns size of note
func (note *eventNote) size() uint {
	if note.Note == SameNote || note.Length == SameLength {
		return 1
	}

	return 2
}

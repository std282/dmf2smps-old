package smpsns

type note byte

const (
	// NoteC = C note
	NoteC note = iota + 0x81
	// NoteCs = C# note
	NoteCs
	// NoteD = D note
	NoteD
	// NoteDs = D# note
	NoteDs
	// NoteE = E note
	NoteE
	// NoteF = F note
	NoteF
	// NoteFs = F# note
	NoteFs
	// NoteG = G note
	NoteG
	// NoteGs = G# note
	NoteGs
	// NoteA = A note
	NoteA
	// NoteAs = A# note
	NoteAs
	// NoteB = B note
	NoteB
)

// GetNote returns byte of a note
func GetNote(nt note, oct int) byte {
	if oct > 7 || (nt == NoteB && oct == 7) {
		logger.Fatal("error: attempted to create SMPS note B-7 or higher")
	}

	return byte(nt) + byte(12*oct)
}

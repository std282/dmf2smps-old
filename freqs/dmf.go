package freqs

// DMF2SMPSNoteFM returns SMPS equivalent of DMF note for FM.
func DMF2SMPSNoteFM(dmfNote, dmfOctave int16) byte {
	switch dmfNote {
	case -1:
		logError.Fatal(
			"empty notes are not supported for conversion",
		)

	case 100:
		return 0x80
	}

	octaveShift := 12 * int(dmfOctave)
	noteShift := int(dmfNote % 12)
	note := byte(0x80 + noteShift + octaveShift)
	switch note {
	case 0xE0, 0xE1:
		logWarn.Print(
			"notes B-7/C-8 are not supported for FM conversion; will treat them as rest notes",
		)

		return 0x80
	}

	return byte(noteShift + octaveShift)
}

// DMF2SMPSNoteDispPSG returns SMPS equivalent of DMF note for PSG.
func DMF2SMPSNoteDispPSG(dmfNote, dmfOctave int16) (note byte, disp int8) {
	// Separately declare and initialize to allow recursion
	var getFreq func(int) int
	getFreq = func(n int) int {
		freq := smpsPSGtableS3[n]

		switch freq {
		case -1:
			return getFreq(n-12) / 2

		case 0x3FF:
			return 0x3F8 // lowest A

		default:
			return freq
		}
	}

	pos := int((dmfNote % 12) + (dmfOctave * 12))
	freq := 0
	switch pos {
	case 95: // B-7
		freq = getFreq(95-12) / 2

	case 96: // C-8
		freq = 0

	default:
		freq = getFreq(pos)
	}

	note, disp = approxPeriodPSG(freq)
	return
}

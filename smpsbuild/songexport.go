package smpsbuild

/*
	This file contains helper functions for exporting SMPS song into a file.
*/

import "io"

// determineChannels determines amount of active FM and PSG channels in song.
//
// When the song is created, all channels are attached to song.stopPattern.
// This function iterates each chip's channels in reverse order:
// FM6-FM1 + DAC and PSG3-PSG1.
//
// If it finds out that channel is not attached to stopPattern, iteration finishes,
// otherwise channels one-by-one excluded from song.
func (song *Song) determineChannels() {
	// FM channels
	song.fmAmount = 7
	for i := len(song.fmHeaders) - 1; i >= 0; i-- {
		initPattern := song.fmHeaders[i].boundPattern
		initAddress := song.fmHeaders[i].dataPointer

		if initPattern != song.stopPattern {
			break
		} else {
			initPattern.removeRef(initAddress)
			song.fmAmount--
		}
	}

	// DAC channel
	{
		initPattern := song.dacHeader.boundPattern
		initAddress := song.dacHeader.dataPointer

		if (song.fmAmount == 1) && (initPattern == song.stopPattern) {
			initPattern.removeRef(initAddress)
			song.fmAmount = 0
		}
	}

	// PSG channels
	song.psgAmount = 3
	for i := len(song.psgHeaders) - 1; i >= 0; i-- {
		initPattern := song.psgHeaders[i].boundPattern
		initAddress := song.psgHeaders[i].dataPointer

		if initPattern != song.stopPattern {
			break
		} else {
			initPattern.removeRef(initAddress)
			song.psgAmount--
		}
	}

	// If detected FM6, then disable DAC
	if song.fmAmount == 7 {
		song.SetInitialPattern(song.stopPattern, DAC)
	}
}

// settleReferences figures out pattern positions in song and values of each
// address (pointer) used in header and coordination flags.
//
// It goes throughtout the whole song, computing position of each element.
// Elements that has associated pointers in the song, use updateRef(pos) function
// to tell those pointers their position.
func (song *Song) settleReferences() {
	var calcPos uint

	// here go SMPS and channel headers
	calcPos += 6 + 4*uint(song.fmAmount) + 5*uint(song.psgAmount)

	// here goes stop pattern
	if song.stopPattern.isReferenced() {
		song.stopPattern.updateRef(calcPos)
		calcPos++
	}

	// here go the voices
	song.voicePtr.evaluate(calcPos)
	for node := song.voices.Front(); node != nil; node = node.Next() {
		calcPos += 25 // size of SMPS voice
	}

	// here go the patterns
	for node := song.noteData.Front(); node != nil; node = node.Next() {
		pat := node.Value.(*Pattern)
		pat.updateRef(calcPos)
		calcPos += pat.eventsSize
	}
}

// removeUnusedPatterns removes patterns that are not referenced anywhere
// throughout the song.
func (song *Song) removeUnusedPatterns() {
	node := song.noteData.Front()
	for node != nil {
		if pat := node.Value.(*Pattern); !pat.isReferenced() {
			del := node
			node = node.Next()
			song.noteData.Remove(del)
		} else {
			node = node.Next()
		}
	}
}

// export writes resolved song to writer.
func (song *Song) export(w io.Writer) {
	// Main header
	song.voicePtr.represent(w)
	w.Write([]byte{
		byte(song.fmAmount),
		byte(song.psgAmount),
		byte(song.tempoDiv),
		byte(song.tempoMod),
	})

	if song.fmAmount > 0 {
		// DAC header
		song.dacHeader.represent(w)

		// FM headers
		for i := 0; i < int(song.fmAmount)-1; i++ {
			song.fmHeaders[i].represent(w)
		}
	}

	// PSG headers
	for i := 0; i < int(song.psgAmount); i++ {
		song.psgHeaders[i].represent(w)
	}

	// Stop pattern
	if song.stopPattern.isReferenced() {
		song.stopPattern.represent(w)
	}

	// FM voices
	for node := song.voices.Front(); node != nil; node = node.Next() {
		voice := node.Value.(*Voice)
		voice.export(w)
	}

	// Patterns
	for node := song.noteData.Front(); node != nil; node = node.Next() {
		pat := node.Value.(*Pattern)
		pat.represent(w)
	}
}

package smpsbuild

import "io"

func (song *Song) determineChannels() {
	song.fmAmount = 7
	for i := len(song.fmHeaders) - 1; i >= 0; i-- {
		if song.fmHeaders[i].boundPattern != song.stopPattern {
			break
		} else {
			song.fmAmount--
		}
	}

	song.psgAmount = 3
	for i := len(song.psgHeaders) - 1; i >= 0; i-- {
		if song.psgHeaders[i].boundPattern != song.stopPattern {
			break
		} else {
			song.psgAmount--
		}
	}

	// If detected FM6, then disable DAC
	if song.fmAmount == 7 {
		song.SetInitialPattern(song.stopPattern, DAC)
	}
}

func (song *Song) settleReferences() {
	var calcPos uint
	// here go SMPS and channel headers
	calcPos += 6 + 4*uint(song.fmAmount) + 5*uint(song.psgAmount)

	// here goes stop pattern
	if song.stopPattern.isReferenced() {
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
		calcPos += pat.size()
	}
}

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

func (song *Song) export(w io.Writer) {
	song.voicePtr.represent(w)
	w.Write([]byte{
		byte(song.fmAmount),
		byte(song.psgAmount),
		byte(song.tempoDiv),
		byte(song.tempoMod),
	})

	song.dacHeader.dataPointer.represent(w)
	for i := uint8(0); i < song.fmAmount; i++ {
		song.fmHeaders[i].represent(w)
	}

	for i := uint8(0); i < song.psgAmount; i++ {
		song.psgHeaders[i].represent(w)
	}

	if song.stopPattern.isReferenced() {
		song.stopPattern.represent(w)
	}

	for node := song.voices.Front(); node != nil; node = node.Next() {
		voice := node.Value.(*Voice)
		voice.export(w)
	}

	for node := song.noteData.Front(); node != nil; node = node.Next() {
		pat := node.Value.(*Pattern)
		pat.represent(w)
	}
}

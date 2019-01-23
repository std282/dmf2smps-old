package freqs

import "math"

var smpsFMtable = [12]int{
	//0x25E, // B
	0x284, // C
	0x2AB, // C#
	0x2D3, // D
	0x2FE, // D#
	0x32D, // E
	0x35C, // F
	0x38F, // F#
	0x3C5, // G
	0x3FF, // G#
	0x43C, // A
	0x47C, // A#
	0x4BC, // B
}

// getEquFreqFM returns equivalent matched FM frequency
func getEquFreqFM(note byte) int {
	var octave uint
	note -= 0x81
	for note >= 12 {
		note -= 12
		octave++
	}

	return smpsFMtable[note] << octave
}

func approxFreqFM(freq int) (note byte, disp int8) {

	// Marginal frequency - maximum frequency that is representable in one FM
	// octave range.
	marginFreq := func() int {
		freqB := float64(smpsFMtable[11])
		quarterToneMult := math.Pow(2, 0.5/12)
		return int(freqB * quarterToneMult)
	}()

	// Computing pre-multiplied frequency
	noteOctaveShift := 0
	for freq > marginFreq {
		freq /= 2
		noteOctaveShift += 12
	}

	// Looking for note whose frequency is the closest to pre-multiplied
	minDisp := 65536
	noteSemitoneShift := 0
	for noteVal, noteFreq := range smpsFMtable {
		disp := freq - noteFreq
		if abs(disp) < abs(minDisp) {
			minDisp = disp
			noteSemitoneShift = noteVal
		}
	}

	// Following switch-case is really unlikely to happen
	warnAboutOverflow := func(got, set int) {
		logWarn.Printf(
			"FM frequency approximation resulted in note displacement overflow; need %d, in fact set %d",
			got,
			set,
		)
	}
	switch {
	case minDisp < -128:
		warnAboutOverflow(minDisp, -128)
		minDisp = -128

	case minDisp > 127:
		warnAboutOverflow(minDisp, 127)
		minDisp = 127
	}

	note = byte(0x81 + noteSemitoneShift + noteOctaveShift)
	disp = int8(minDisp)
	return
}

// OptimizeFM finds optimal match for note and its displacement. That is, it
// finds such note that the displacement is as small as possible.
func OptimizeFM(inNote byte, inDisp int) (outNote byte, outDisp int8) {
	freq := getEquFreqFM(inNote) + inDisp
	outNote, outDisp = approxFreqFM(freq)
	return
}

// DisplaceFM displaces optimized note to specified amount of semitones
func DisplaceFM(inNote byte, inDisp int8, semitones float64) (outNote byte, outDisp int8) {
	freq := func() int {
		orig := getEquFreqFM(inNote) + int(inDisp)
		mult := math.Pow(2, semitones/12)
		resFreq := float64(orig) * mult
		return int(resFreq)
	}()

	outNote, outDisp = approxFreqFM(freq)
	return
}

// DisplaceRawFM displaces unoptimized note to specified amount of semitones
func DisplaceRawFM(inNote byte, inDisp int, semitones float64) (outNote byte, outDisp int8) {
	freq := func() int {
		orig := getEquFreqFM(inNote) + inDisp
		mult := math.Pow(2, semitones/12)
		resFreq := float64(orig) * mult
		return int(resFreq)
	}()

	outNote, outDisp = approxFreqFM(freq)
	return
}

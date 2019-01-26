package freqs

import "math"

var smpsPSGtableS1 = [12*8 - 1]int{
	0x356, 0x326, 0x2F9, 0x2CE, 0x2A5, 0x280, 0x25C, 0x23A, 0x21A, 0x1FB, 0x1DF, 0x1C4,
	0x1AB, 0x193, 0x17D, 0x167, 0x153, 0x140, 0x12E, 0x11D, 0x10D, 0xFE, 0xEF, 0xE2,
	0xD6, 0xC9, 0xBE, 0xB4, 0xA9, 0xA0, 0x97, 0x8F, 0x87, 0x7F, 0x78, 0x71,
	0x6B, 0x65, 0x5F, 0x5A, 0x55, 0x50, 0x4B, 0x47, 0x43, 0x40, 0x3C, 0x39,
	0x36, 0x33, 0x30, 0x2D, 0x2B, 0x28, 0x26, 0x24, 0x22, 0x20, 0x1F, 0x1D,
	0x1B, 0x1A, 0x18, 0x17, 0x16, 0x15, 0x13, 0x12, 0x11, 0, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
}

// I shall never need this.

// var smpsPSGtableMy = [12*8 - 1]int{
// 	0x356, 0x326, 0x2F9, 0x2CE, 0x2A5, 0x280, 0x25C, 0x23A, 0x21A, 0x1FB, 0x1DF, 0x1C4,
// 	0x1AB, 0x193, 0x17D, 0x167, 0x153, 0x140, 0x12E, 0x11D, 0x10D, 0xFE, 0xEF, 0xE2,
// 	0xD6, 0xC9, 0xBE, 0xB4, 0xA9, 0xA0, 0x97, 0x8F, 0x87, 0x7F, 0x78, 0x71,
// 	0x6B, 0x65, 0x5F, 0x5A, 0x55, 0x50, 0x4B, 0x47, 0x43, 0x40, 0x3C, 0x39,
// 	0x36, 0x33, 0x30, 0x2D, 0x2B, 0x28, 0x26, 0x24, 0x22, 0x20, 0x1F, 0x1D,
// 	0x1B, 0x1A, 0x18, 0x17, 0x16, 0x15, 0x13, 0x12, 0x11, 0, -1, -1,
// 	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
// 	-1, -1, -1, -1, -1, -1, -1, -1, 0x3F8, 0x3BF, 0x389,
// }

var smpsPSGtableS3 = [12*8 - 1]int{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, 0x3F8, 0x3BF, 0x389,
	0x356, 0x326, 0x2F9, 0x2CE, 0x2A5, 0x280, 0x25C, 0x23A, 0x21A, 0x1FB, 0x1DF, 0x1C4,
	0x1AB, 0x193, 0x17D, 0x167, 0x153, 0x140, 0x12E, 0x11D, 0x10D, 0xFE, 0xEF, 0xE2,
	0xD6, 0xC9, 0xBE, 0xB4, 0xA9, 0xA0, 0x97, 0x8F, 0x87, 0x7F, 0x78, 0x71,
	0x6B, 0x65, 0x5F, 0x5A, 0x55, 0x50, 0x4B, 0x47, 0x43, 0x40, 0x3C, 0x39,
	0x36, 0x33, 0x30, 0x2D, 0x2B, 0x28, 0x26, 0x24, 0x22, 0x20, 0x1F, 0x1D,
	0x1B, 0x1A, 0x18, 0x17, 0x16, 0x15, 0x13, 0x12, 0x11, 0, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
}

var smpsPSGtable = smpsPSGtableS1[:]

type psgTable int

const (
	// PSGTableFromSonic1 speaks for itself
	PSGTableFromSonic1 psgTable = iota
	// PSGTableFromSonic3 speaks for itself
	PSGTableFromSonic3
)

// SetPSGTable sets global PSG table. PSG notes will be adjusted and calculated
// according to the table selected.
func SetPSGTable(table psgTable) {
	switch table {
	case PSGTableFromSonic1:
		smpsPSGtable = smpsPSGtableS1[:]

	case PSGTableFromSonic3:
		smpsPSGtable = smpsPSGtableS3[:]
	}
}

func getPeriodPSG(note byte) int {
	period := smpsPSGtable[note-0x81]
	if period == -1 {
		logError.Fatalf(
			"PSG note $%02X has invalid period",
			note,
		)
	}

	return period
}

func approxPeriodPSG(period int) (note byte, disp int8) {
	if period > 0x3FF {
		logWarn.Printf(
			"impossible PSG period (%d); will be set to 0x3FF",
			period,
		)
		period = 0x3FF
	}

	minShift, minDiff := 0, 65536
	for shift, notePeriod := range smpsPSGtable {
		if notePeriod == -1 {
			continue
		}

		diff := period - notePeriod
		if abs(diff) < abs(minDiff) {
			minDiff = diff
			minShift = shift
		}
	}

	warnAboutOverflow := func(got, set int) {
		logWarn.Printf(
			"PSG period approximation resulted in note displacement overflow; need %d, in fact set %d",
			got,
			set,
		)
	}

	switch {
	case minDiff < -128:
		warnAboutOverflow(minDiff, -128)
		minDiff = -128

	case minDiff > 127:
		warnAboutOverflow(minDiff, 127)
		minDiff = 127
	}

	note = byte(minShift + 0x81)
	disp = int8(minDiff)
	return
}

// OptimizePSG finds optimal match for note and its displacement. That is, it
// finds such note that the displacement is as small as possible.
func OptimizePSG(oldNote byte, oldDisp int) (newNote byte, newDisp int8) {
	period := getPeriodPSG(oldNote) + oldDisp
	newNote, newDisp = approxPeriodPSG(period)
	return
}

// DisplacePSG displaces optimized note to specified amount of semitones
func DisplacePSG(inNote byte, inDisp int8, semitones float64) (oldNote byte, oldDisp int8) {
	period := func() int {
		orig := getPeriodPSG(inNote) + int(inDisp)
		div := math.Pow(2, semitones/12)
		resPeriod := float64(orig) / div
		return int(resPeriod)
	}()

	oldNote, oldDisp = approxPeriodPSG(period)
	return
}

// DisplaceRawPSG displaces unoptimized note to specified amount of semitones
func DisplaceRawPSG(inNote byte, inDisp int, semitones float64) (oldNote byte, oldDisp int8) {
	period := func() int {
		orig := getPeriodPSG(inNote) + int(inDisp)
		div := math.Pow(2, semitones/12)
		resPeriod := float64(orig) / div
		return int(resPeriod)
	}()

	oldNote, oldDisp = approxPeriodPSG(period)
	return
}

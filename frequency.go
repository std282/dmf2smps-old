package main

var fmFreqArray [12*8 - 1]int

var psgFreqArray = [12*6 - 2]int{
	0x356, 0x326, 0x2F9, 0x2CE, 0x2A5, 0x280, 0x25C, 0x23A, 0x21A, 0x1FB, 0x1DF, 0x1C4,
	0x1AB, 0x193, 0x17D, 0x167, 0x153, 0x140, 0x12E, 0x11D, 0x10D, 0xFE, 0xEF, 0xE2,
	0xD6, 0xC9, 0xBE, 0xB4, 0xA9, 0xA0, 0x97, 0x8F, 0x87, 0x7F, 0x78, 0x71,
	0x6B, 0x65, 0x5F, 0x5A, 0x55, 0x50, 0x4B, 0x47, 0x43, 0x40, 0x3C, 0x39,
	0x36, 0x33, 0x30, 0x2D, 0x2B, 0x28, 0x26, 0x24, 0x22, 0x20, 0x1F, 0x1D,
	0x1B, 0x1A, 0x18, 0x17, 0x16, 0x15, 0x13, 0x12, 0x11, 0,
}

func init() {
	fmFreqArray[0] = 0x284
	fmFreqArray[1] = 0x2AB
	fmFreqArray[2] = 0x2D3
	fmFreqArray[3] = 0x2FE
	fmFreqArray[4] = 0x32D
	fmFreqArray[5] = 0x35C
	fmFreqArray[6] = 0x38F
	fmFreqArray[7] = 0x3C5
	fmFreqArray[8] = 0x3FF
	fmFreqArray[9] = 0x43C
	fmFreqArray[10] = 0x47C
	fmFreqArray[11] = 0x4C0

	for i := 12; i < len(fmFreqArray); i++ {
		fmFreqArray[i] = fmFreqArray[i-12] + 606
	}
}

// GetFMFreq returns corrected FM frequency
func GetFMFreq(smpsNote byte) int {
	return fmFreqArray[smpsNote-0x81]
}

// GetPSGFreq returns PSG frequency (period)
func GetPSGFreq(smpsNote byte) int {
	return psgFreqArray[smpsNote-0x81]
}

// ComputeBestFreqDispFM computes best frequency displacement for FM
func ComputeBestFreqDispFM(smpsNote byte, initDisp int) (resNote byte, resDisp int) {
	resFreq := fmFreqArray[smpsNote-0x81] + initDisp
	intAbs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	resNote = byte(0)
	resDisp = 65536 // definitely more than max FM frequency
	for i := range fmFreqArray {
		if disp := fmFreqArray[i] - resFreq; intAbs(disp) < intAbs(resDisp) {
			resNote = byte(i) + 0x81
			resDisp = disp
		}
	}

	return
}

// ComputeBestFreqDispPSG computes best frequency displacement for PSG
func ComputeBestFreqDispPSG(smpsNote byte, disp int) (resNote byte, resDisp int) {
	resFreq := psgFreqArray[smpsNote-0x81] + resDisp
	intAbs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	resNote = byte(0)
	resDisp = 65536 // definitely more than max PSG frequency/period
	for i := range psgFreqArray {
		if disp := psgFreqArray[i] - resFreq; intAbs(disp) < resDisp {
			resNote = byte(i) + 0x81
			resDisp = disp
		}
	}

	return
}

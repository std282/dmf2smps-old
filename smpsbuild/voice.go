package smpsbuild

import "io"

// Voice represents one SMPS voice
type Voice struct {
	FB, ALG int

	MULT, DT, RS, AR, DR, SR, SL, RR, TL [4]int
}

func (voice *Voice) export(w io.Writer) {
	var voiceArr [25]byte

	voice.ALG &= 7 // 3 bits
	voice.FB &= 7  // 3 bits
	voiceArr[0] = byte(voice.ALG | (voice.FB << 3))
	op := []int{0, 2, 1, 3} // operator sequence
	for i, j := range op {
		voice.AR[i] &= 31   // 5 bits
		voice.DR[i] &= 31   // 5 bits
		voice.DT[i] &= 7    // 3 bits
		voice.MULT[i] &= 15 // 4 bits
		voice.RR[i] &= 15   // 4 bits
		voice.RS[i] &= 3    // 2 bits
		voice.SL[i] &= 15   // 4 bits
		voice.SR[i] &= 127  // 7 bits
		voice.TL[i] &= 127  // 7 bits

		voiceArr[1+j] = byte(voice.MULT[i] | (voice.DT[i] << 4))
		voiceArr[5+j] = byte(voice.AR[i] | (voice.RS[i] << 6))
		voiceArr[9+j] = byte(voice.DR[i])
		voiceArr[13+j] = byte(voice.SR[i])
		voiceArr[17+j] = byte(voice.RR[i] | (voice.SL[i] << 4))
		voiceArr[21+j] = byte(voice.TL[i])
	}

	w.Write(voiceArr[:])
}

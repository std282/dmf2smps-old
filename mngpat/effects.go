package mngpat

// SetModulation adds modulation to pattern
func (mpat *ManagedPattern) SetModulation(delay, stretch, factor, magnitude int) {
	sameMod := delay == mpat.modDelay ||
		stretch == mpat.modStretch ||
		factor == mpat.modFactor ||
		magnitude == mpat.modMagnitude

	if sameMod {
		mpat.innerPattern.ResumeModulation()
	} else {
		mpat.SetModulation(delay, stretch, factor, magnitude)

		mpat.modDelay = delay
		mpat.modStretch = stretch
		mpat.modFactor = factor
		mpat.modMagnitude = magnitude
	}

	mpat.modActive = true
}

// StopModulation stops modulation in pattern
func (mpat *ManagedPattern) StopModulation() {
	mpat.innerPattern.StopModulation()
	mpat.modActive = false
}

// AlterVolume alters volume of pattern
func (mpat *ManagedPattern) AlterVolume(value int) {
	if mpat.isPSG {
		mpat.innerPattern.AlterVolumePSG(value)
	} else {
		mpat.innerPattern.AlterVolumeFM(value)
	}
}

// SetVoice sets current voice of pattern
func (mpat *ManagedPattern) SetVoice(voice int) {
	if mpat.voice == voice {
		if mpat.isPSG {
			mpat.innerPattern.SetPSGVoice(voice)
		} else {
			mpat.innerPattern.SetFMVoice(voice)
		}
	}
}

// SetFineTune sets fine tuning of pattern
func (mpat *ManagedPattern) SetFineTune(value int) {
	if value != mpat.freqDisp {
		mpat.innerPattern.SetFineTune(value)
		mpat.freqDisp = value
	}
}

// SetPanning sets panning of pattern
func (mpat *ManagedPattern) SetPanning(side byte) {
	if side != mpat.panning {
		mpat.innerPattern.SetPan(side)
		mpat.panning = side
	}
}

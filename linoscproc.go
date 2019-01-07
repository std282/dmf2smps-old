package main

import "math"

// LinOscProc describes sum of linear and oscillating processes
type LinOscProc struct {
	// Linear
	magnitude float64 // amount of units added to result every frame

	// Oscillating
	amplitude float64 // amplitude of oscillation
	period    int     // period of oscillation, frames
	phase     int     // phase of oscillation, frames
	rectMask  uint8   // rectification mask

	// Limits
	lower int // lower bound of result
	upper int // upper bound of result

	linstate float64 // linear state
	oscstate float64 // oscillator state
}

// NewLOP returns new LOP, ready to use
func NewLOP() *LinOscProc {
	return &LinOscProc{rectMask: 3}
}

// SetLinear sets linear part of LOP
func (lop *LinOscProc) SetLinear(mag float64) {
	lop.magnitude = mag
}

// SetLinearFrac sets linear part of LOP (fraction)
func (lop *LinOscProc) SetLinearFrac(magN int, magD int) {
	lop.magnitude = float64(magN) / float64(magD)
}

// SetOscillator sets oscillatior part of LOP
func (lop *LinOscProc) SetOscillator(amp float64, per int) {
	lop.amplitude = amp
	lop.period = per
	lop.phase = 0
}

// SetLimits sets limits of modelling
func (lop *LinOscProc) SetLimits(min, max int) {
	lop.lower = min
	lop.upper = max
}

// SetOscRect sets rectification of oscillation process
func (lop *LinOscProc) SetOscRect(up, down bool) {
	switch {
	case up && down || !up && !down:
		lop.rectMask = 0x03

	case up:
		lop.rectMask = 0x01

	case down:
		lop.rectMask = 0x02
	}
}

// PokeResult forcefully changes current state
func (lop *LinOscProc) PokeResult(res float64) {
	lop.linstate = res
}

// PeekResult returns current state
func (lop *LinOscProc) PeekResult() int {
	res := int(math.Round(lop.linstate + lop.oscstate))
	switch {
	case res < lop.lower:
		res = lop.lower

	case res > lop.upper:
		res = lop.upper
	}

	return res
}

// Update iterates process one frame further, changing the state
func (lop *LinOscProc) Update() int {
	lop.linstate += lop.magnitude

	if lop.period > 0 {
		phase := 2 * math.Pi / float64(lop.period) * float64(lop.phase)
		lop.oscstate = lop.amplitude * math.Sin(phase)

		// Rectification
		upBlock := (lop.rectMask & 0x01) == 0
		downBlock := (lop.rectMask & 0x02) == 0
		if (upBlock && lop.oscstate > 0) || (downBlock && lop.oscstate < 0) {
			lop.oscstate = 0
		}

		lop.phase++
		if lop.phase == lop.period {
			lop.phase = 0
		}
	}

	return lop.PeekResult()
}

// Freeze stops process, making it stationary
func (lop *LinOscProc) Freeze() {
	lop.amplitude = 0
	lop.period = 0
	lop.phase = 0

	lop.magnitude = 0
}

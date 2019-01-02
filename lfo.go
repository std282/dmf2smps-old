package main

import "math"

// LFO represents low frequency oscillator to emulate tremolo and vibrato
type LFO struct {
	period   int
	min, max int

	t int
}

// NewLFO creates initialized LFO with specified period, min, max
func NewLFO(period, min, max int) *LFO {
	lfo := LFO{
		period: period,
		min:    min,
		max:    max,
		t:      0,
	}
	return &lfo
}

// Init initializes LFO with specified period, min, max
func (lfo *LFO) Init(period, min, max int) {
	lfo.period = period
	lfo.min = min
	lfo.max = max
	lfo.t = 0
}

// Iterate iterates oscillating process of LFO
func (lfo *LFO) Iterate() int {
	if !lfo.IsActive() {
		logger.Fatal("error: cannot iterate on inactive LFO")
	}

	const pi = math.Pi
	period := float64(lfo.period)
	min := float64(lfo.min)
	max := float64(lfo.max)
	t := float64(lfo.t)

	// Compute sine and scale it so that it's in range [min : max]
	res := math.Sin(2 * pi / period * t)
	res *= (max - min) / 2
	res += (max + min) / 2

	lfo.t++
	return int(math.Round(res))
}

// EquiModulation returns settings for SMPS modulation flag (F0) that can be used
// to approximate LFO on channel frequency
func (lfo *LFO) EquiModulation() (time, amp, trig int) {
	twoAmp := lfo.max - lfo.min
	if twoAmp > lfo.period {
		trig = lfo.period - 1
		amp = twoAmp / lfo.period
		time = 1
	} else if twoAmp < lfo.period {
		trig = twoAmp
		amp = 1
		time = (lfo.period - 1) / twoAmp
	} else {
		trig = lfo.period - 1
		amp = 1
		time = 1
	}

	return
}

// Deactivate makes LFO inactive. When inactive, LFO cannot be iterated
func (lfo *LFO) Deactivate() {
	lfo.period = 0
}

// IsActive returns true if LFO is active; false otherwise
func (lfo *LFO) IsActive() bool {
	return lfo.period != 0
}

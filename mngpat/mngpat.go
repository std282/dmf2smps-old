package mngpat

import (
	"github.com/std282/dmf2smps/smpsbuild"
)

// ManagedPattern is a clever pattern wrapper. It places data without spending
// extra place (cleverly puts notes, prevents placing useless effects, etc).
//
// Effectively it's just pattern with state.
type ManagedPattern struct {
	innerPattern *smpsbuild.Pattern
	isPSG        bool

	// Notes
	note       byte
	noteLength int8

	// Modulation
	modDelay     int
	modStretch   int
	modFactor    int
	modMagnitude int
	modActive    bool

	// Frequency displacement
	freqDisp int

	// Panning
	panning byte

	// Voices
	voice int
}

// NewFM returns new managed pattern which is built specifically for FM channels.
func NewFM() *ManagedPattern {
	return &ManagedPattern{
		innerPattern: smpsbuild.NewPattern(),
		panning:      smpsbuild.PanCenter,
	}
}

// NewPSG returns new managed pattern which is built specifically for PSG channels.
func NewPSG() *ManagedPattern {
	return &ManagedPattern{
		innerPattern: smpsbuild.NewPattern(),
		isPSG:        true,
	}
}

// NewFrom returns new managed pattern which state is copied from prototype
// pattern.
func NewFrom(proto *ManagedPattern) *ManagedPattern {
	copy := *proto
	copy.innerPattern = smpsbuild.NewPattern()
	return &copy
}

// Pattern returns underlying pattern pointer. You are expected to use it for
// finer settings.
func (mpat *ManagedPattern) Pattern() *smpsbuild.Pattern {
	pat := mpat.innerPattern
	mpat.innerPattern = nil
	return pat
}

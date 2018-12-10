package smpsns

import (
	"github.com/std282/dmf2smps/smpsbuild"
)

type panMode byte

const (
	// Left panning
	Left panMode = 0x80
	// Right panning
	Right panMode = 0x40
	// Center panning
	Center panMode = 0xC0
)

// SetPan sets panning. Only for FM channels
func SetPan(pat *smpsbuild.Pattern, pan panMode) {
	pat.AddBytes(0xE0, byte(pan))
}

// SetChanFreqDisp sets channel frequency displacement.
// disp must not be out of range -128...127
func SetChanFreqDisp(pat *smpsbuild.Pattern, disp int) {
	pat.AddBytes(0xE1, byte(disp))
}

// Return returns from subpattern
func Return(pat *smpsbuild.Pattern) {
	pat.AddBytes(0xE3)
}

// SetTempoDivLocal sets tempo divider for current track
func SetTempoDivLocal(pat *smpsbuild.Pattern, tempoDiv int) {
	pat.AddBytes(0xE5, byte(tempoDiv))
}

// AddVolume changes volume of the current track
func AddVolume(pat *smpsbuild.Pattern, volDisp int) {
	pat.AddBytes(0xE6, byte(volDisp))
}

// TODO

// SetModulation adds modulation flag
func SetModulation(pat *smpsbuild.Pattern, delay int, mult int, period int) {
	pat.AddBytes(0xF0, byte(delay), byte(mult), byte(period))
}

// TODO

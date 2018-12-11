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

type noiseMode byte

const (
	// Periodic noise mode
	Periodic noiseMode = 0xE3
	// Random noise mode
	Random noiseMode = 0xE7
)

type loopPriority byte

const (
	// Low loop priority
	Low loopPriority = 2
	// Middle loop priority
	Middle loopPriority = 1
	// High loop priority
	High loopPriority = 0
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

// AlterVolumeFM changes volume of the current FM track
func AlterVolumeFM(pat *smpsbuild.Pattern, volDisp int) {
	pat.AddBytes(0xE6, byte(volDisp))
}

// NoAttack prevents next note from attacking
func NoAttack(pat *smpsbuild.Pattern) {
	pat.AddBytes(0xE7)
}

// CutNote cuts note after several frames
func CutNote(pat *smpsbuild.Pattern, fill int) {
	pat.AddBytes(0xE8, byte(fill))
}

// AlterPitch alters pitch by semitones
func AlterPitch(pat *smpsbuild.Pattern, alter int) {
	pat.AddBytes(0xE9, byte(alter))
}

// SetTempoModGlobal sets tempo modifier globally
func SetTempoModGlobal(pat *smpsbuild.Pattern, mod int) {
	pat.AddBytes(0xEA, byte(mod))
}

// SetTempoDivGlobal sets tempo divider globally
func SetTempoDivGlobal(pat *smpsbuild.Pattern, div int) {
	pat.AddBytes(0xEB, byte(div))
}

// AlterVolumePSG changes volume of the current PSG track
func AlterVolumePSG(pat *smpsbuild.Pattern, volDisp int) {
	pat.AddBytes(0xEC, byte(volDisp))
}

// SetFMVoice sets current FM voice
func SetFMVoice(pat *smpsbuild.Pattern, voice int) {
	pat.AddBytes(0xEF, byte(voice))
}

// SetModulation adds modulation (vibrato)
func SetModulation(pat *smpsbuild.Pattern, delay int, mult int, period int) {
	pat.AddBytes(0xF0, byte(delay), byte(mult), byte(period))
}

// EnableMod enables last disabled modulation
func EnableMod(pat *smpsbuild.Pattern) {
	pat.AddBytes(0xF1)
}

// Stop makes entire song stop
func Stop(pat *smpsbuild.Pattern) {
	pat.AddBytes(0xF2)
}

// SetNoise sets PSG3 noise mode
func SetNoise(pat *smpsbuild.Pattern, mode noiseMode) {
	pat.AddBytes(0xF3, byte(mode))
}

// DisableMod deactivates active modulation
func DisableMod(pat *smpsbuild.Pattern) {
	pat.AddBytes(0xF4)
}

// SetPSGEnvelope sets envelope for PSG channel
func SetPSGEnvelope(pat *smpsbuild.Pattern, env int) {
	pat.AddBytes(0xF5, byte(env))
}

// Jump continues song execution from other location
func Jump(pat *smpsbuild.Pattern, loc smpsbuild.RelAddress) {
	pat.AddBytes(0xF6)
	pat.AddAddress(&loc)
}

// Loop makes section from 'addr' to current location repeat several 'times'
func Loop(pat *smpsbuild.Pattern, addr smpsbuild.RelAddress, times int, prior loopPriority) {
	pat.AddBytes(0xF7, byte(prior), byte(times))
	pat.AddAddress(&addr)
}

// Call continues song execution from other location, but it will return later
func Call(pat *smpsbuild.Pattern, loc smpsbuild.RelAddress) {
	pat.AddBytes(0xF8)
	pat.AddAddress(&loc)
}

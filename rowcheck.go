package main

import (
	"fmt"

	"github.com/std282/dmf2smps/dmfns"
	"github.com/std282/dmf2smps/dmfparse"
)

var chNames = []string{
	"FM1",
	"FM2",
	"FM3",
	"FM4",
	"FM5",
	"FM6",
	"STD1",
	"STD2",
	"STD3",
	"STD4",
}

// CheckRow checks row for being valid. If it's invalid, it prints warning
func CheckRow(row dmfparse.Row, chNo, patNo, rowNo int) {
	for i := 0; i < row.EffectsAmount(); i++ {
		if fx := row.Effects[i].Type; fx != -1 && !IsEffectValid(fx) {
			logger.Println(
				fmt.Sprint(
					"warning: inconvertable effect (%02.X) at channel %v, ",
					"pattern #%X, row #%d; it will be ignored",
				),
				fx,
				chNames[chNo],
				patNo,
				rowNo,
			)
		}
	}
}

// IsEmpty returns true if row can be safely ignored
func IsEmpty(row dmfparse.Row) bool {
	noteEmpty := row.Note == -1
	octvEmpty := row.Octave == -1
	instEmpty := row.InstNum == -1
	voluEmpty := row.Volume == -1
	fxEmpty := true
	for i := 0; i < row.EffectsAmount(); i++ {
		fxEmpty = fxEmpty && !IsEffectValid(row.Effects[i].Type)
	}

	return noteEmpty && octvEmpty && instEmpty && voluEmpty && fxEmpty
}

// IsEffectValid returns true if effect is valid and convertable; false otherwise
func IsEffectValid(fxType int16) bool {
	switch fxType {
	case dmfns.ArpSpeed, dmfns.Arpeggio, dmfns.BreakPattern, dmfns.EnableDAC,
		dmfns.FineTune, dmfns.GoToPattern, dmfns.NoiseMode, dmfns.NoteCut,
		dmfns.NoteDelay, dmfns.Panning, dmfns.PortaDown, dmfns.PortaDownFix,
		dmfns.PortaNote, dmfns.PortaUp, dmfns.PortaUpFix, dmfns.Retrigger,
		dmfns.SampleBank, dmfns.SetSpeed1, dmfns.SetSpeed2, dmfns.Tremolo,
		dmfns.Vibrato, dmfns.VibratoDepth, dmfns.VibratoMode, dmfns.VolSlide,
		dmfns.VolSlidePortaNote, dmfns.VolSlideVibrato:
		return true

	default:
		return false
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/std282/dmf2smps/dmfparse"
	"github.com/std282/dmf2smps/smpsbuild"
)

var globalParams struct {
	dmfSong *dmfparse.Song
}

func convert(dts ConvDetails) smpsbuild.Song {
	smps := smpsbuild.NewSong()

	dmfFile, err := os.OpenFile(dts.FileName, os.O_RDONLY, 0)
	if err != nil {
		logger.Fatal("error: cannot open file specified in conversion settings")
	}

	globalParams.dmfSong = dmfparse.NewSongParse(dmfFile)

	// TODO: tempo conversion
	// TODO: events conversion

	return smps
}

func convertInstrument(inst *dmfparse.InstrumentFM, pos int) smpsbuild.Voice {
	var voice smpsbuild.Voice
	voice.ALG = int(inst.ALG)
	voice.FB = int(inst.FB)

	if inst.LFO != 0 || inst.LFO2 != 0 {
		logger.Printf(
			fmt.Sprint(
				"warning: FM instrument #%v, name \"%v\" ",
				"has active LFO; it will be ignored in SMPS",
			),
			pos,
			inst.Name(),
		)
	}

	for i := 0; i < 4; i++ {
		voice.AR[i] = int(inst.AR[i])
		voice.DR[i] = int(inst.DR[i])
		voice.DT[i] = int(inst.DT[i])
		voice.MULT[i] = int(inst.MULT[i])
		voice.RR[i] = int(inst.RR[i])
		voice.RS[i] = int(inst.RS[i])
		voice.RS[i] = int(inst.RS[i])
		voice.SL[i] = int(inst.SL[i])
		voice.TL[i] = int(inst.TL[i])

		if inst.AM[i] != 0 {
			logger.Printf(
				fmt.Sprint(
					"warning: FM instrument #%v, name \"%v\", operator %v ",
					"has active AM; it will be ignored in SMPS",
				),
				pos,
				inst.Name(),
				i,
			)
		}

		if inst.SSG[i] != 0 {
			logger.Printf(
				fmt.Sprint(
					"warning: FM instrument #%v, name \"%v\", operator %v ",
					"has active SSG; it will be ignored in SMPS",
				),
				pos,
				inst.Name(),
				i,
			)
		}
	}

	return voice
}

func (ngn *FMEngine) init(rowFramesPtr *int) {
	ngn.state.lastNote = 0x80
	ngn.state.instrument = -1
	ngn.state.arp.length = 1
	ngn.rowLen = rowFramesPtr
}

func (ngn *FMEngine) fetch(row dmfparse.Row) {
	if IsEmpty(row) {
		ngn.acmFrames += *ngn.rowLen
	} else {
		if ngn.acmFrames != 0 {
			releaseAccum(&ngn.buf, ngn.acmFrames)
		}
	}
	// TODO
}

func (ngn *FMEngine) yield() []byte {
	return flushBuffer(&ngn.buf)
}

func (ngn *FMEngine) getState() FMState {
	return ngn.state
}

func (ngn *DACEngine) init() {
	// TODO
}

func (ngn *DACEngine) fetch() {
	// TODO
}

func (ngn *DACEngine) yield() []byte {
	return flushBuffer(&ngn.buf)
}

func (ngn *DACEngine) getState() DACState {
	return ngn.state
}

func (ngn *PSGEngine) init() {
	// TODO
}

func (ngn *PSGEngine) fetch() {
	// TODO
}

func (ngn *PSGEngine) yield() []byte {
	return flushBuffer(&ngn.buf)
}

func (ngn *PSGEngine) getState() PSGState {
	return ngn.state
}

func (ngn *NoiseEngine) init() {
	// TODO
}

func (ngn *NoiseEngine) fetch() {
	// TODO
}

func (ngn *NoiseEngine) yield() []byte {
	return flushBuffer(&ngn.buf)
}

func (ngn *NoiseEngine) getState() NoiseState {
	return ngn.state
}

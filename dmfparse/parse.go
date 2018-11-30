package dmfparse

import (
	"io"
	"log"
	"strconv"
)

func localError(description string) {
	log.Panic("dmfparse: " + description)
}

func localErrorf(format string, args ...interface{}) {
	log.Panicf(
		"dmfparse: "+format, args...)
}

// Parse parses DMF song which is being read from reader
func (song *Song) Parse(r io.Reader) (err error) {
	pr := panicReader{reader: r, eof: false}
	defer recover()

	// Parsing preamble and version. We quit if preamble fails
	if preamble := pr.ReadString(16); preamble != ".DelekDefleMask." {
		localError("file is not a DefleMask module")
	}
	if version := pr.Read8(); version != 24 {
		localErrorf("unsupported version (%v), must be 24", version)
	}

	// Actual parsing
	song.parse(pr) // headers

	for i := range song.Matrix { // pattern matrix
		for j := range song.Matrix[i] {
			song.Matrix[i][j] = int(pr.Read8())
		}
	}

	instAmount := int(pr.Read8()) // instruments
	for i := 0; i < instAmount; i++ {
		instName := pr.ReadPascalString()

		switch instType := pr.Read8(); instType {
		case 1: // FM
			inst := InstrumentFM{name: instName}
			inst.parse(pr)
			song.Instruments = append(song.Instruments, inst)

		case 0: // STD
			inst := InstrumentSTD{name: instName}
			inst.parse(pr)
			song.Instruments = append(song.Instruments, inst)

		default:
			localErrorf("instrument #%v has invalid type", i)
		}
	}

	wtAmount := int(pr.Read8()) // wavetables
	for i := 0; i < wtAmount; i++ {
		wtSize := int(pr.Read32())
		pr.Skip(wtSize * 4)
	}

	for i := range song.Channels { // channels
		song.Channels[i].parse(pr)
	}

	sampAmount := int(pr.Read8()) // samples
	song.Samples = make([]Sample, sampAmount)
	for i := range song.Samples {
		song.Samples[i].parse(pr)
	}

	return
}

func (song *Song) parse(pr panicReader) {
	song.Name = pr.ReadPascalString()
	song.Author = pr.ReadPascalString()
	pr.Skip(2)

	song.TimeBase, song.TickTime1, song.TickTime2 =
		int(pr.Read8()), int(pr.Read8()), int(pr.Read8())

	framesMode, hasCustom := pr.Read8(), pr.Read8()
	custom := pr.ReadString(3)
	if hasCustom == 0 {
		if framesMode == 1 {
			song.FramesPerSecond = 60
		} else {
			song.FramesPerSecond = 50
		}
	} else {
		var err error
		song.FramesPerSecond, err = strconv.Atoi(custom)
		if err != nil {
			localError(err.Error())
		}
	}

	patternSize := pr.Read32()
	matrixSize := pr.Read8()

	for i := range song.Matrix {
		song.Matrix[i] = make([]int, matrixSize)
	}

	for i := range song.Channels {
		song.Channels[i].Rows = make([][]Row, matrixSize)
		for j := range song.Channels[i].Rows {
			song.Channels[i].Rows[j] = make([]Row, patternSize)
			for k := range song.Channels[i].Rows[j] {
				song.Channels[i].Rows[j][k].parent = &song.Channels[i]
			}
		}
	}
}

func (inst *InstrumentFM) parse(pr panicReader) {
	// TODO
}

func (inst *InstrumentSTD) parse(pr panicReader) {
	// TODO
}

func (ch *Channel) parse(pr panicReader) {
	// TODO
}

func (row *Row) parse(pr panicReader) {
	// TODO
}

func (fx *Effect) parse(pr panicReader) {
	// TODO
}

func (samp *Sample) parse(pr panicReader) {
	// TODO
}

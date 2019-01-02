package dmfparse

import (
	"io"
	"strconv"
)

// NewSong returns new DMF song
func NewSong() *Song {
	return &Song{}
}

// NewSongParse returns DMF song parsed from r
func NewSongParse(r io.Reader) *Song {
	song := NewSong()
	song.Parse(r)
	return song
}

// Parse parses DMF song which is being read from reader
func (song *Song) Parse(r io.Reader) {
	pr := panicReader{reader: r, eof: false}

	// Parsing preamble and version. We quit if preamble fails
	if preamble := pr.ReadString(16); preamble != ".DelekDefleMask." {
		logger.Fatal("error: file is not a DefleMask module")
	}
	if version := pr.Read8(); version != 24 {
		logger.Printf("error: unsupported version (%v), must be 24", version)
		logger.Fatal(
			"hint: try to open this module in the latest version of ",
			"DefleMask and resave it",
		)
	}
	if system := pr.Read8(); system != 2 {
		logger.Fatalf(
			"error: specified system (%v) is not supported, use SEGA Genesis",
			systemName[system],
		)
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
			logger.Fatalf("error: instrument #%v has invalid type", i)
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
			logger.Fatal("error: cannot parse custom FPS")
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
	inst.ALG = pr.Read8()
	inst.FB = pr.Read8()
	inst.LFO = pr.Read8()
	inst.LFO2 = pr.Read8()
	opseq := []int{1, 3, 2, 4}
	for _, op := range opseq {
		i := op - 1
		inst.AM[i] = pr.Read8()
		inst.AR[i] = pr.Read8()
		inst.DR[i] = pr.Read8()
		inst.MULT[i] = pr.Read8()
		inst.RR[i] = pr.Read8()
		inst.SL[i] = pr.Read8()
		inst.TL[i] = pr.Read8()
		inst.DT2[i] = pr.Read8()
		inst.RS[i] = pr.Read8()
		inst.DT[i] = pr.Read8()
		inst.D2R[i] = pr.Read8()
		inst.SSG[i] = pr.Read8()
	}
}

func (inst *InstrumentSTD) parse(pr panicReader) {
	if vollen := pr.Read8(); vollen > 0 {
		inst.VolumeEnv = make([]int32, vollen)
		pr.ReadAny(inst.VolumeEnv)
		inst.VolumeLoop = int(pr.Read8s())
	}

	if arplen := pr.Read8(); arplen > 0 {
		inst.ArpeggioEnv = make([]int32, arplen)
		pr.ReadAny(inst.ArpeggioEnv)
		inst.ArpeggioLoop = int(pr.Read8s())
	}
	inst.ArpeggioMode = int(pr.Read8())

	if noilen := pr.Read8(); noilen > 0 {
		inst.NoiseEnv = make([]int32, noilen)
		pr.ReadAny(inst.NoiseEnv)
		inst.NoiseLoop = int(pr.Read8s())
	}

	if wtlen := pr.Read8(); wtlen > 0 {
		pr.Skip(int(wtlen))
	}
}

func (ch *Channel) parse(pr panicReader) {
	ch.effectsAmount = int(pr.Read8())
	for i := range ch.Rows {
		for j := range ch.Rows[i] {
			ch.Rows[i][j].Effects = make([]Effect, ch.effectsAmount)
			ch.Rows[i][j].parse(pr)
		}
	}
}

func (row *Row) parse(pr panicReader) {
	row.Note = pr.Read16()
	row.Octave = pr.Read16()
	row.Volume = pr.Read16()
	for i := range row.Effects {
		row.Effects[i].parse(pr)
	}
	row.InstNum = pr.Read16()
}

func (fx *Effect) parse(pr panicReader) {
	fx.Type = pr.Read16()
	fx.Byte = pr.Read16()
}

func (samp *Sample) parse(pr panicReader) {
	size := int(pr.Read32())
	samp.Name = pr.ReadPascalString()
	pr.Skip(4 + 2*size)
}

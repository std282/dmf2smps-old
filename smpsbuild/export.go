package smpsbuild

import (
	"encoding/binary"
	"io"
	"log"
)

// Anything that could be exported in SMPS
type exportable interface {
	sizeOf() uint       // tells the size in bytes when exported
	export(w io.Writer) // exports to byte writer
}

func (song *Song) sizeOf() uint {
	retSize := song.headerSize()
	for i := range song.voices {
		retSize += song.voices[i].sizeOf()
	}

	for elem := song.data.Front(); elem != nil; elem = elem.Next() {
		if val, ok := elem.Value.(Pattern); ok {
			retSize += val.sizeOf()
		} else {
			log.Fatal("smpsbuild: non-exportable entity in song.data")
		}
	}

	return retSize
}

func (*Voice) sizeOf() uint {
	return 25
}

func (*absAddress) sizeOf() uint {
	return 2
}

func (*relAddress) sizeOf() uint {
	return 2
}

func (chunk *byteChunk) sizeOf() uint {
	return uint(chunk.buf.Len())
}

func (pat *Pattern) sizeOf() uint {
	var accumSize uint
	for el := pat.events.Front(); el != nil; el = el.Next() {
		if exp, ok := el.Value.(exportable); ok {
			accumSize += exp.sizeOf()
		} else {
			log.Fatal("smpsbuild: non-exportable entity in pattern")
		}
	}

	return accumSize
}

func (song *Song) export(w io.Writer) {
	voicePtr := absAddress{pointer: uint16(song.headerSize())}
	voicePtr.export(w)

	w.Write([]byte{
		byte(song.channelsFM),
		byte(song.channelsPSG),
		byte(song.tempoDivider),
		byte(song.tempoModifier),
	})

	song.offsetDAC.export(w)
	w.Write([]byte{0x00, 0x00})

	for i := 0; i < song.channelsFM-1; i++ {
		song.offsetFM[i].export(w)

		w.Write([]byte{
			byte(song.pitchFM[i]),
			byte(song.volumeFM[i]),
		})
	}

	for i := 0; i < song.channelsPSG; i++ {
		song.offsetPSG[i].export(w)

		w.Write([]byte{
			byte(song.pitchPSG[i]),
			byte(song.volumePSG[i]),
			byte(song.voicePSG[i]),
		})
	}

	for i := range song.voices {
		song.voices[i].export(w)
	}

	for el := song.data.Front(); el != nil; el = el.Next() {
		if pat, ok := el.Value.(*Pattern); ok {
			pat.export(w)
		} else {
			log.Fatal("smpsbuild: non-exportable entity in song.data")
		}
	}
}

func (voice *Voice) export(w io.Writer) {
	var voiceRepr [25]byte

	// 0 0 (3 bits of FB) (3 bits of ALG)
	voiceRepr[0] = byte((voice.ALG & 0x07) | ((voice.FB & 0x07) << 3))

	for i := 0; i < 4; i++ {
		// (4 bits of DT) (4 bits of MULT)
		voiceRepr[1+i] = byte((voice.MULT[i] & 15) | ((voice.DT[i] & 15) << 4))

		// (2 bits of RS) 0 (5 bits of AR)
		voiceRepr[5+i] = byte((voice.AR[i] & 31) | ((voice.DT[i] & 3) << 6))

		// DR
		voiceRepr[9+i] = byte(voice.DR[i])

		// SR
		voiceRepr[13+i] = byte(voice.SR[i])

		// (4 bits of SL) (4 bits of RR)
		voiceRepr[17+i] = byte((voice.RR[i] & 15) | ((voice.SL[i] & 15) << 4))

		// TL
		voiceRepr[21+i] = byte(voice.TL[i])
	}

	_, err := w.Write(voiceRepr[:])
	if err != nil {
		log.Fatal(err.Error())
	}

	return
}

func (addr *absAddress) export(w io.Writer) {
	err := binary.Write(w, binary.BigEndian, addr.pointer)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (addr *relAddress) export(w io.Writer) {
	err := binary.Write(w, binary.BigEndian, addr.pointer)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (chunk *byteChunk) export(w io.Writer) {
	var bytebuf [256]byte

	var err error
	var n int
	for err == nil {
		n, err = chunk.buf.Read(bytebuf[:])
		if n > 0 {
			w.Write(bytebuf[0:n])
		}
	}

	if err != io.EOF {
		log.Fatal(err.Error())
	}
}

func (pat *Pattern) export(w io.Writer) {
	for el := pat.events.Front(); el != nil; el = el.Next() {
		if exp, ok := el.Value.(exportable); ok {
			exp.export(w)
		} else {
			log.Fatal("smpsbuild: non-exportable entity in pattern")
		}
	}
}

// Resolves every address in song
func (song *Song) resolveAddresses() {
	count := song.headerSize() + uint(len(song.voices))*25

	for el := song.data.Front(); el != nil; el = el.Next() {
		pat, ok := el.Value.(*Pattern)
		if !ok {
			log.Fatalf(
				"smpsbuild: song contains something that is not a pattern: %T",
				el.Value,
			)
		}

		pat.setRelAddrPos(count)
		pat.visit(count)
		size := pat.sizeOf()
		count += size
	}
}

package smpsbuild

import (
	"encoding/binary"
	"errors"
	"io"
)

/*
	- Exportable interface declaration:
		Size() uint
		Export(io.Writer)
	- Exportable implementation for:
		Song
		Voice
		AbsAddress
		RelAddress
		Chunk
*/

// Exportable describes a thing that can be exported in binary format
type Exportable interface {
	Size() (uint, error)
	Export(w io.Writer) error
}

// Size returns size of an exported song
func (song *Song) Size() (uint, error) {
	retSize := song.headerSize()
	for i := range song.voices {
		size, _ := song.voices[i].Size()
		retSize += size
	}

	for elem := song.data.Front(); elem != nil; elem = elem.Next() {
		if val, ok := elem.Value.(Pattern); ok {
			size, _ := val.Size()
			retSize += size
		} else {
			return retSize, errors.New("unexportable thing in song.data")
		}
	}

	return retSize, nil
}

// Size returns size of an exported voice
func (*Voice) Size() (uint, error) {
	return 25, nil
}

// Size returns size of an exported absolute address
func (*AbsAddress) Size() (uint, error) {
	return 2, nil
}

// Size returns size of an exported relative address
func (*RelAddress) Size() (uint, error) {
	return 2, nil
}

// Size returns size of an exported chunk of bytes
func (chunk *Chunk) Size() (uint, error) {
	return uint(chunk.buf.Len()), nil
}

// Size returns size of pattern
func (pat *Pattern) Size() (uint, error) {
	var accumSize uint
	for el := pat.events.Front(); el != nil; el = el.Next() {
		if exp, ok := el.Value.(Exportable); ok {
			size, _ := exp.Size()
			accumSize += size
		} else {
			return accumSize, errors.New("non-exportable entity in pattern")
		}
	}

	return accumSize, nil
}

// Export exports song
func (song *Song) Export(w io.Writer) error {
	count := song.headerSize()
	voicePtr := AbsAddress{pointer: uint16(count)}
	voicePtr.Export(w)

	w.Write([]byte{
		byte(song.channelsFM),
		byte(song.channelsPSG),
		byte(song.TempoDivider),
		byte(song.TempoModifier),
	})

	for i := 0; i < song.channelsFM; i++ {
		song.offsetFM[i].Export(w)

		w.Write([]byte{
			byte(song.pitchFM[i]),
			byte(song.volumeFM[i]),
		})
	}

	for i := 0; i < song.channelsPSG; i++ {
		song.offsetPSG[i].Export(w)

		w.Write([]byte{
			byte(song.pitchPSG[i]),
			byte(song.volumePSG[i]),
			byte(song.voicePSG[i]),
		})
	}

	for i := range song.voices {
		song.voices[i].Export(w)
	}

	for el := song.data.Front(); el != nil; el = el.Next() {
		if pat, ok := el.Value.(Pattern); ok {
			err := pat.Export(w)
			if err != nil {
				return err
			}
		} else {
			return errors.New("non-exportable entity in song.data")
		}
	}

	return nil
}

// Export exports voice
func (voice *Voice) Export(w io.Writer) (err error) {
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

	_, err = w.Write(voiceRepr[:])
	return
}

// Export exports absolute address
func (addr *AbsAddress) Export(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, addr.pointer)
}

// Export exports relative address
func (addr *RelAddress) Export(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, addr.pointer)
}

// Export exports chunk of bytes
func (chunk *Chunk) Export(w io.Writer) error {
	var bytebuf [256]byte

	var err error
	var n int
	for err == nil {
		n, err = chunk.buf.Read(bytebuf[:])
		if n > 0 {
			w.Write(bytebuf[0 : n-1])
		}
	}

	if err == io.EOF {
		return nil
	}

	return err
}

// Export exports pattern
func (pat *Pattern) Export(w io.Writer) error {
	for el := pat.events.Front(); el != nil; el = el.Next() {
		if exp, ok := el.Value.(Exportable); ok {
			err := exp.Export(w)
			if err != nil {
				return err
			}
		} else {
			return errors.New("non-exportable entity in pattern")
		}
	}

	return nil
}

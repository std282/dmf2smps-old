package smpsbuild

import (
	"fmt"
	"strings"
)

type BinaryStringWriter struct {
	writtenBytes uint
	builder      strings.Builder
}

func (bsw *BinaryStringWriter) Write(p []byte) (n int, err error) {
	for _, b := range p {
		bsw.builder.WriteString(fmt.Sprintf("%02X ", b))
		bsw.writtenBytes++
		if (bsw.writtenBytes % 16) == 0 {
			bsw.builder.WriteString("\n")
		}
	}

	return len(p), nil
}

func (bsw *BinaryStringWriter) Stop() {
	bsw.builder.WriteString("\n\n")
}

func (bsw *BinaryStringWriter) Release() string {
	return bsw.builder.String()
}

// func TestPattern(t *testing.T) {
// 	bsw := new(BinaryStringWriter)
// 	bsw.Write([]byte{
// 		0x00, 0x00,
// 		0x06, 0x03,
// 		0x01, 0x00,
// 	})

// 	pat := NewPattern()
// 	pat.PlaceNoteClever(0x81, 16)
// 	pat.PlaceNoteClever(0x82, 16)
// 	pat.PlaceNoteClever(0x81, 8)
// 	pat.PlaceNoteClever(0x81, 8)
// 	pat.PlaceNoteClever(0x82, 16)
// 	pat.PlaceLoop(2, 5)

// 	pat2 := NewPattern()
// 	pat2.PlaceNoteClever(0x85, 16)
// 	pat2.PlaceNoteClever(0x82, 16)
// 	pat2.PlaceNoteClever(0x81, 16)
// 	pat2.PlaceNoteClever(0x82, 16)
// 	pat2.PlaceJump(pat)

// 	pat.updateRef(6)
// 	pat2.updateRef(6 + pat.eventsSize)

// 	pat.represent(bsw)
// 	pat2.represent(bsw)
// 	bsw.Stop()

// 	fmt.Print(bsw.Release())
// }

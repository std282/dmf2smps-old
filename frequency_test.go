package main

import (
	"fmt"
	"testing"
)

func TestFrequency(test *testing.T) {
	for disp := 100; disp < 2000; disp += 100 {
		testSingleNote(0x81+4+12*1, disp)
	}
}

func testSingleNote(note byte, disp int) {
	resNote, resDisp := ComputeBestFreqDispFM(note, disp)
	fmt.Printf("$%02X (%d) = $%02X (%d)\n", note, disp, resNote, resDisp)
}

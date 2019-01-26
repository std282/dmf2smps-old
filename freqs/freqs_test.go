package freqs

import (
	"fmt"
	"testing"
)

func TestFM(t *testing.T) {
	note, disp := byte(0xB7), 0
	for i := 0; i < 30; i++ {
		newNote, newDisp := OptimizeFM(note, disp-20)
		note = newNote
		disp = int(newDisp)
		fmt.Printf("$%02X, %d\n", newNote, newDisp)
	}
}

func TestPSG(t *testing.T) {
	note, disp := byte(0xB7), 0
	for i := 0; i < 20; i++ {
		newNote, newDisp := OptimizePSG(note, disp+20)
		note = newNote
		disp = int(newDisp)
		fmt.Printf("$%02X, %d\n", newNote, newDisp)
	}
}

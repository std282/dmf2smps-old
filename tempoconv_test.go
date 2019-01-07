package main

import (
	"fmt"
	"testing"
)

func TestTempoConv(test *testing.T) {
	for fps := 20; fps < 70; fps++ {
		tempo := GetOptimalTempo(fps)
		fmt.Printf("Tested for FPS = %d, tempo = %02X\n\n", fps, tempo)
	}
}

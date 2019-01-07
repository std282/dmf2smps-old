package main

import (
	"fmt"
	"testing"
)

func TestLOP(test *testing.T) {
	lop := NewLOP()
	lop.SetLimits(0, 2000)
	lop.PokeResult(0)

	fmt.Println("Testing magnitude...")
	lop.SetLinearFrac(131, 7)
	for i := 0; i < 10; i++ {
		state := lop.Update()
		fmt.Printf("Step %d: %d\n", i+1, state)
	}
	fmt.Println()

	lop.Freeze()

	fmt.Println("Testing oscillations...")
	lop.SetOscillator(20, 5)
	for i := 0; i < 10; i++ {
		state := lop.Update()
		fmt.Printf("Step %d: %d\n", i+1, state)
	}
	fmt.Println()

	lop.Freeze()

	fmt.Println("Testing both...")
	lop.SetLinearFrac(-10, 1)
	lop.SetOscillator(5, 10)
	for i := 0; i < 20; i++ {
		state := lop.Update()
		fmt.Printf("Step %d: %d\n", i+1, state)
	}
}

package main

import (
	"fmt"
	"testing"
)

func TestRowCheck(test *testing.T) {
	fxTest := func(fx int16) {
		ok := IsEffectValid(fx)
		fmt.Printf("Effect %02X: ", fx)
		if ok {
			fmt.Println("success")
		} else {
			fmt.Println("failure")
		}
	}

	for i := int16(0); i < 0xFF; i++ {
		fxTest(i)
	}
}

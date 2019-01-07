package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/std282/dmf2smps/dmfparse"
)

func TestAnalyze(test *testing.T) {
	dmfFile, _ := os.OpenFile("gu.dmf", os.O_RDONLY, 0)
	defer dmfFile.Close()
	dmfEnc, _ := zlib.NewReader(dmfFile)
	dmf := dmfparse.NewSongParse(dmfEnc)

	json := Analyze(dmf, "gu.dmf")
	config, _ := os.Create("config_test.json")
	io.WriteString(config, json)
	defer config.Close()
}

func TestRetrieve(test *testing.T) {
	jsonFile, _ := os.OpenFile("config_test.json", os.O_RDONLY, 0)
	defer jsonFile.Close()

	dts := Retrieve(jsonFile)
	fmt.Printf("%d PSG mappings, %d DAC mappings\n", len(dts.PSGMap), len(dts.DACMap))

	for _, val := range dts.PSGMap {
		fmt.Printf("STD #%d: \"%v\" - %02X\n", val.InstNumber, val.Name, val.Envelope)
	}

	for _, val := range dts.DACMap {
		fmt.Printf(
			"Sample @%v/%d: \"%v\" - %02X\n",
			val.Note,
			val.Bank,
			val.Name,
			val.Sample,
		)
	}
}

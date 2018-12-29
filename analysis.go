package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/std282/dmf2smps/dmfns"
	"github.com/std282/dmf2smps/dmfparse"
)

// Returns JSON string with analysis results
func analyze(dmf *dmfparse.Song, path string) string {
	var dts ConvDetails
	dts.FileName = path

	dts.PreferFM6 = false
	dts.PreferPSG3 = false
	dts.DecayVibrato = false
	dts.RestartAfterEnd = true
	dts.ExtendedPSG = false

	dts.getInstData(dmf)
	dts.getSampleData(dmf)

	jsonStr, err := json.MarshalIndent(dts, "", "    ")
	if err != nil {
		log.Fatalf("dmf2smps: error: json marshalling failed (%v)", err.Error())
	}

	return string(jsonStr)
}

// Returns JSON-able array of sample mapping objects
func (dts *ConvDetails) getSampleData(dmf *dmfparse.Song) {
	curNote := dmfns.NoteC
	curBank := 1
	for i := range dmf.Samples {
		dts.addDACEntry(curNote, curBank, dmf.Samples[i].Name)

		switch curNote {
		case dmfns.NoteC:
			curNote = dmfns.NoteCs

		case dmfns.NoteB:
			curBank++
			curNote = dmfns.NoteC

		default:
			curNote++
		}
	}
}

// Returns JSON-able array of STD instrument mapping objects
func (dts *ConvDetails) getInstData(dmf *dmfparse.Song) {
	for i := range dmf.Instruments {
		if dmf.Instruments[i].Type() == dmfparse.STD {
			dts.addPSGEntry(i, dmf.Instruments[i].Name())
		}
	}
}

func retrieve(jsonReader io.Reader) ConvDetails {
	jsonRaw, err := ioutil.ReadAll(jsonReader)
	if err != nil {
		log.Fatal("dmf2smps: error: unable to read configuration file")
	}

	var dts ConvDetails
	err = json.Unmarshal(jsonRaw, &dts)
	if err != nil {
		log.Fatal("dmf2smps: error: unable to unmarshall JSON conversion parameters")
	}

	// We need to convert HEX strings to numbers in DAC map
	for i := range dts.DACMap {
		dts.DACMap[i].Sample = parseNumberInMapping(
			dts.DACMap[i].Sample,
			i,
			dts.DACMap[i].Name,
			true, // for DAC
		)
	}

	for i := range dts.PSGMap {
		dts.PSGMap[i].Envelope = parseNumberInMapping(
			dts.PSGMap[i].Envelope,
			i,
			dts.PSGMap[i].Name,
			false, // for PSG
		)
	}

	return dts
}

func parseNumberInMapping(smth interface{}, pos int, name string, isDAC bool) interface{} {
	var place, field string
	if isDAC {
		place = "DAC mapping array"
		field = "dacSample"
	} else {
		place = "PSG mapping array"
		field = "psgEnvelope"
	}

	switch smth.(type) {
	case float64:
		return int(smth.(float64))

	case string:
		val, err := strconv.ParseInt(smth.(string), 16, 8)
		if err != nil {
			log.Printf(
				fmt.Sprint(
					"dmf2smps: warning: in %v, could not parse 8-bit hex number at ",
					"position %v, name \"%v\"; will treat this sample as ignored",
				),
				place,
				pos,
				name,
			)

			return nil
		}

		return val
	}

	if smth != nil {
		log.Fatalf(
			fmt.Sprint(
				"dmf2smps: error: in %v, field \"%v\" is of invalid type at ",
				"position %v, name \"%v\"",
			),
			place,
			field,
			pos,
			name,
		)
	}

	return nil
}

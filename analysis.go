package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/std282/dmf2smps/dmfns"
	"github.com/std282/dmf2smps/dmfparse"
)

// Analyze returns JSON string with analysis results
func Analyze(dmf *dmfparse.Song, path string) string {
	var dts ConvDetails
	dts.FileName = path

	dts.PreferFM6 = false
	dts.PreferPSG3 = false
	dts.DecayVibrato = false
	dts.RestartAfterEnd = true
	dts.ExtendedPSG = false

	dts.GetInstData(dmf)
	dts.GetSampleData(dmf)

	jsonStr, err := json.MarshalIndent(dts, "", "    ")
	if err != nil {
		logger.Fatalf("error: json marshalling failed (%v)", err.Error())
	}

	return string(jsonStr)
}

// GetSampleData returns JSON-able array of sample mapping objects
func (dts *ConvDetails) GetSampleData(dmf *dmfparse.Song) {
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

// GetInstData returns JSON-able array of STD instrument mapping objects
func (dts *ConvDetails) GetInstData(dmf *dmfparse.Song) {
	for i := range dmf.Instruments {
		if dmf.Instruments[i].Type() == dmfparse.STD {
			dts.AddPSGEntry(i, dmf.Instruments[i].Name())
		}
	}
}

// Retrieve unmarshals JSON conversion structure from reader
func Retrieve(jsonReader io.Reader) ConvDetails {
	jsonRaw, err := ioutil.ReadAll(jsonReader)
	if err != nil {
		logger.Fatal("error: unable to read configuration file")
	}

	var dts ConvDetails
	err = json.Unmarshal(jsonRaw, &dts)
	if err != nil {
		logger.Fatal("error: unable to unmarshall JSON conversion parameters")
	}

	// We need to convert HEX strings to numbers in DAC map
	for i := range dts.DACMap {
		dts.DACMap[i].Sample = ParseNumberInMapping(
			dts.DACMap[i].Sample,
			i,
			dts.DACMap[i].Name,
			true, // for DAC
		)
	}

	for i := range dts.PSGMap {
		dts.PSGMap[i].Envelope = ParseNumberInMapping(
			dts.PSGMap[i].Envelope,
			i,
			dts.PSGMap[i].Name,
			false, // for PSG
		)
	}

	return dts
}

// ParseNumberInMapping parses number in mapping of unspecified JSON type.
// It can be either string or number or null.
// String = hexadecimal number
// Int = just number
// Null = ignored or decayed
func ParseNumberInMapping(smth interface{}, pos int, name string, isDAC bool) interface{} {
	var place, field, object, nullState string
	if isDAC {
		place = "DAC mapping array"
		field = "dacSample"
		object = "sample"
		nullState = "ignored"
	} else {
		place = "PSG mapping array"
		field = "psgEnvelope"
		object = "envelope"
		nullState = "decayed to volume alterations"
	}

	switch smth.(type) {
	case float64:
		return int(smth.(float64))

	case string:
		val, err := strconv.ParseInt(smth.(string), 16, 8)
		if err != nil {
			logger.Printf(
				"warning: in %v, could not parse 8-bit hex number at position %v, name \"%v\"; will treat this %v as %v",
				place,
				pos,
				name,
				object,
				nullState,
			)

			return nil
		}

		return val
	}

	if smth != nil {
		logger.Fatalf(
			"error: in %v, field \"%v\" is of invalid type at position %v, name \"%v\"",
			place,
			field,
			pos,
			name,
		)
	}

	return nil
}

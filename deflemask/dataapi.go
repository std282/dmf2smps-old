package deflemask

// GetRow returns specified row.
func (dmf *Module) GetRow(channel, pattern, row int) *Row {
	return &dmf.songData[channel][pattern][row]
}

// GetPattern returns specified pattern.
func (dmf *Module) GetPattern(channel, pattern int) []Row {
	return dmf.songData[channel][pattern]
}

// GetInstrumentType returns type of specified instrument.
func (dmf *Module) GetInstrumentType(inst int) int {
	switch dmf.instruments[inst].(type) {
	case *InstrumentFM:
		return FM

	case *InstrumentSTD:
		return STD

	default: // unlikely
		logErr.Panicf("instrument #%d is of unknown type", inst)
	}

	return -1 // never reached
}

// GetInstrumentFM asserts that instrument at specified position is of FM type
// and returns pointer to FM instrument.
func (dmf *Module) GetInstrumentFM(inst int) *InstrumentFM {
	return dmf.instruments[inst].(*InstrumentFM)
}

// GetInstrumentSTD asserts that instrument at specified position is of STD type
// and returns pointer to STD instrument.
func (dmf *Module) GetInstrumentSTD(inst int) *InstrumentSTD {
	return dmf.instruments[inst].(*InstrumentSTD)
}

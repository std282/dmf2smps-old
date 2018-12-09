package smpsbuild

/*
	struct
		Song
		Voice
		AbsAddress
		RelAddress
		Chunk

	interface
		Exportable
		Address

*/

import (
	"bytes"
	"container/list"
)

// Song represents SMPS song
type Song struct {
	// Main header data
	voices        []Voice
	TempoDivider  int
	TempoModifier int
	channelsFM    int
	channelsPSG   int

	// Channel header data
	offsetFM  []AbsAddress
	volumeFM  []int
	pitchFM   []int
	offsetPSG []AbsAddress
	volumePSG []int
	pitchPSG  []int
	voicePSG  []int

	// Note data
	data list.List
}

func (song *Song) headerSize() uint {
	// 6 bytes for main header
	// 4 * FMChannels for FM channel header
	// 6 * PSGChannels for PSG channel header
	return uint(6 + song.channelsFM*4 + song.channelsPSG*6)
}

// Voice represents SMPS FM voice
type Voice struct {
	FB, ALG int

	MULT, DT, AR, RS, DR, SR, RR, SL, TL [4]int
}

// Address describes reference to a chunk of bytes
type Address interface {
	Set(from uint)         // sets address
	Refer(ref Addressable) // address will be set later
}

// AbsAddress represents deferring evaluation absolute address
type AbsAddress struct {
	referencesTo Addressable
	pointer      uint16
}

// Refer makes address refer to a chunk of bytes which position is unknown
func (addr *AbsAddress) Refer(ref Addressable) {
	addr.referencesTo = ref // make address refer to chunk
	ref.Notify(addr)        // notify the chunk about being referenced
}

// Set sets address to location "from"
func (addr *AbsAddress) Set(from uint) {
	addr.pointer = uint16(from)
}

// RelAddress represents deferring evaluation relative address
type RelAddress struct {
	refto    Addressable
	location uint
	pointer  int16
}

// Set sets pointer to location "from"
func (addr *RelAddress) Set(from uint) {
	addr.pointer = int16(from - addr.location)
}

// Refer makes address refer to a chunk of bytes
func (addr *RelAddress) Refer(ref Addressable) {
	addr.refto = ref
	ref.Notify(addr)
}

// Chunk represents and exportable chunk of raw bytes
type Chunk struct {
	buf bytes.Buffer
}

// Addressable desribes anything that could be referenced
type Addressable interface {
	Notify(byWhom Address)
	Visit(curPos uint)
}

// Pattern represents any addressable sequence of SMPS events
type Pattern struct {
	events   list.List
	refdFrom list.List

	lastIsBytes bool
}

// Notify tells pattern that someone referenced him, but doesn't know what
// address it has. It is for pattern to be able to call back the referer and
// tell them the address
func (pat *Pattern) Notify(byWhom Address) {
	pat.refdFrom.PushBack(byWhom)
}

// Visit makes pattern to look at who asked him to tell his address and actually
// set their address
func (pat *Pattern) Visit(curPos uint) {
	for el := pat.refdFrom.Front(); el != nil; el = el.Next() {
		if addr, ok := el.Value.(Address); ok {
			addr.Set(curPos)
		}
	}
}

package smpsbuild

import (
	"container/list"
	"io"
)

// Song represents SMPS song
type Song struct {
	// Main header data
	Voices        []Voice
	TempoDivider  int
	TempoModifier int
	ChannelsFM    int
	ChannelsPSG   int

	// Channel header data
	OffsetFM  []AbsAddress
	VolumeFM  []int
	PitchFM   []int
	OffsetPSG []AbsAddress
	VolumePSG []int
	PitchPSG  []int
	VoicePSG  []int

	// Note data
	data list.List
}

// Exportable describes a thing that can be exported in binary format
type Exportable interface {
	Size() uint
	Export(w io.Writer)
}

// Voice represents SMPS FM voice
type Voice struct {
	FB, ALG int

	MULT, DT, AR, RS, DR, SR, RR, SL, TL [4]int
}

// Size returns size of a voice when exported
func (*Voice) Size() uint {
	return 25
}

// AbsAddress represents deferring evaluation absolute address
type AbsAddress struct {
	refto   *Chunk
	pointer uint16
}

// Refer makes address refer to a chunk of bytes which position is unknown
func (addr *AbsAddress) Refer(chunk *Chunk) {
	addr.refto = chunk           // make address refer to chunk
	chunk.reffrom.PushBack(addr) // notify the chunk about being referenced
}

// Set sets address to location "from"
func (addr *AbsAddress) Set(from uint) {
	addr.pointer = uint16(from)
}

// Size returns a size of address when exported
func (*AbsAddress) Size() uint {
	return 2
}

// RelAddress represents deferring evaluation relative address
type RelAddress struct {
	refto    *Chunk
	location uint
	pointer  int16
}

// Size returns a size of address when exported
func (*RelAddress) Size() uint {
	return 2
}

// Set sets pointer to location "from"
func (addr *RelAddress) Set(from uint) {
	addr.pointer = int16(from - addr.location)
}

// Refer makes address refer to a chunk of bytes
func (addr *RelAddress) Refer(chunk *Chunk) {
	addr.refto = chunk           // make address refer to chunk
	chunk.reffrom.PushBack(addr) // notify the chunk about being referenced
}

// Address describes reference to a chunk of bytes
type Address interface {
	Set(from uint)      // sets address
	Refer(chunk *Chunk) // address will be set later
}

// Chunk represents a byte array
type Chunk struct {
	data    []byte
	reffrom list.List
}

// Size returns size of chunk when exported
func (chunk *Chunk) Size() uint {
	return uint(len(chunk.data))
}

// VisitReferences visits all addresses referenced to a current chunk and sets
// their pointers to a value of counter
func (chunk *Chunk) VisitReferences(counter uint) {
	for elem := chunk.reffrom.Front(); elem != nil; elem = elem.Next() {
		if val, ok := elem.Value.(Address); ok {
			val.Set(counter)
		}
	}
}

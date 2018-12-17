package smpsbuild

import (
	"bytes"
)

// Address describes reference to a chunk of bytes
type Address interface {
	set(from uint)         // sets address
	refer(ref Addressable) // address will be set later
}

// Absolute address
type absAddress struct {
	refTo   Addressable // what it references to
	pointer uint16      // its value
}

func (addr *absAddress) refer(ref Addressable) {
	addr.refTo = ref
	ref.notify(addr)
}

func (addr *absAddress) set(from uint) {
	addr.pointer = uint16(from)
}

// Relative address
type relAddress struct {
	refTo    Addressable // what it references to
	location uint        // where it's located
	pointer  int16       // its value
}

func (addr *relAddress) set(from uint) {
	addr.pointer = int16(from - addr.location)
}

func (addr *relAddress) refer(ref Addressable) {
	addr.refTo = ref
	ref.notify(addr)
}

// A chunk of bytes
type byteChunk struct {
	buf bytes.Buffer // buffer that holds the bytes
}

// Addressable describes anything that could be referenced
type Addressable interface {
	notify(byWhom Address) // take into account about being referenced
	visit(curPos uint)     // tell everyone referenced about position
}

func (pat *Pattern) notify(byWhom Address) {
	pat.refdFrom.PushBack(byWhom)
}

func (pat *Pattern) visit(curPos uint) {
	for el := pat.refdFrom.Front(); el != nil; el = el.Next() {
		if addr, ok := el.Value.(Address); ok {
			addr.set(curPos)
		}
	}

	pat.refdFrom.Init()
}

func (song *Song) headerSize() uint {
	// 6 bytes for main header
	// 4 * FMChannels for FM channel header
	// 6 * PSGChannels for PSG channel header
	return uint(6 + song.channelsFM*4 + song.channelsPSG*6)
}

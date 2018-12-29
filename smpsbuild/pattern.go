package smpsbuild

import "container/list"

// Pattern represents any addressable sequence of SMPS events (like notes,
// coordination flags...)
type Pattern struct {
	events   list.List // chunks, addresses
	refdFrom list.List // addresses that refer to this pattern

	lastIsBytes bool // true, if last event is chunk of bytes
}

// AddBytes adds raw bytes to pattern
func (pat *Pattern) AddBytes(b ...byte) {
	if !pat.lastIsBytes {
		pat.events.PushBack(&byteChunk{})
		pat.lastIsBytes = true
	}

	chunk := pat.events.Back().Value.(*byteChunk)
	chunk.buf.Write(b)
}

// AddAddress adds address to pattern
func (pat *Pattern) AddAddress(addr Address) {
	pat.events.PushBack(addr)
	pat.lastIsBytes = false
}

// CreatePattern creates pattern in the song and returns pointer to it
func (song *Song) CreatePattern() *Pattern {
	pat := Pattern{}
	song.data.PushBack(&pat)
	return &pat
}

// RemovePattern removes the pattern from song
func (song *Song) RemovePattern(pat *Pattern) {
	toDelete := list.New()

	for node := song.data.Front(); node != nil; node = node.Next() {
		nodePat := node.Value.(*Pattern)
		if nodePat == pat {
			nodePat.unbind()
			toDelete.PushBack(node)
		}
	}

	for node := toDelete.Front(); node != nil; node = node.Next() {
		del := node.Value.(*list.Element)
		song.data.Remove(del)
	}
}

// GetAddress returns relative address of the pattern.
//
// We do not need absolute addresses. Those are only used in headers, and they're
// handled separately
func (pat *Pattern) GetAddress() Address {
	addr := relAddress{}
	addr.refer(pat)
	return &addr
}

// Resolve every referenced-to address
func (pat *Pattern) setRelAddrPos(patPos uint) {
	for el := pat.events.Front(); el != nil; el = el.Next() {
		switch val := el.Value.(type) {
		case *byteChunk:
			patPos += val.sizeOf()

		case *relAddress:
			val.location = patPos + 1
			patPos += val.sizeOf()
		}
	}
}

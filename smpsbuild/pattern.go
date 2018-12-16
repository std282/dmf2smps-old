package smpsbuild

// AddBytes adds bytes to pattern
func (pat *Pattern) AddBytes(b ...byte) {
	if !pat.lastIsBytes {
		pat.events.PushBack(&Chunk{})
		pat.lastIsBytes = true
	}

	chunk := pat.events.Back().Value.(*Chunk)
	chunk.buf.Write(b)
}

// AddAddress adds address to pattern
func (pat *Pattern) AddAddress(addr Address) {
	pat.events.PushBack(addr)
	pat.lastIsBytes = false
}

// AddPattern appends one pattern to the song
func (song *Song) AddPattern(pat *Pattern) {
	song.data.PushBack(pat)
}

// CreatePattern creates pattern in the song and returns pointer to it
func (song *Song) CreatePattern() *Pattern {
	pat := Pattern{}
	song.data.PushBack(&pat)
	return &pat
}

// GetAddress returns relative address of the pattern.
//
// We do not need absolute addresses. Those are only used in headers, and they're
// handled separately
func (pat *Pattern) GetAddress() Address {
	addr := RelAddress{}
	addr.Refer(pat)
	return &addr
}

// SetRelAddrPos sets position of all relative addresses in pattern
func (pat *Pattern) SetRelAddrPos(patPos uint) {
	for el := pat.events.Front(); el != nil; el = el.Next() {
		switch val := el.Value.(type) {
		case *Chunk:
			size, _ := val.Size()
			patPos += size

		case *RelAddress:
			size, _ := val.Size()
			val.location = patPos + 1
			patPos += size
		}
	}
}

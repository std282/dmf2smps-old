package smpsbuild

// AddBytes adds bytes to pattern
func (pat *Pattern) AddBytes(b ...byte) {
	if !pat.lastIsBytes {
		pat.events.PushBack(Chunk{})
		pat.lastIsBytes = true
	}

	chunk := pat.events.Back().Value.(Chunk)
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

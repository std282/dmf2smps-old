package smpsbuild

func (chunk *Chunk) AddBytes(bts ...byte) {
	chunk.data = append(chunk.data, bts...)
}

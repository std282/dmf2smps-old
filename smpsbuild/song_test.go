package smpsbuild

// func TestSMPSBuild(t *testing.T) {
// 	song := NewSong()
// 	pat1 := NewPattern()
// 	pat2 := NewPattern()
// 	pat3 := NewPattern()

// 	pat1.PlaceNoteClever(0x81, 16)
// 	pat1.PlaceNoteClever(0x82, 16)
// 	pat1.PlaceNoteClever(0x81, 8)
// 	pat1.PlaceNoteClever(0x81, 8)
// 	pat1.PlaceNoteClever(0x82, 16)
// 	pat1.PlaceJump(pat3)

// 	pat2.PlaceNoteClever(0x81, 16)
// 	pat2.PlaceNoteClever(0x82, 12)
// 	pat2.PlaceNoteClever(0x81, 8)
// 	pat2.PlaceNoteClever(0x82, 4)
// 	pat2.PlaceNoteClever(0x81, 8)
// 	pat2.PlaceNoteClever(0x82, 4)
// 	pat2.PlaceNoteClever(0x82, 4)
// 	pat2.PlaceNoteClever(0x82, 4)
// 	pat2.PlaceNoteClever(0x82, 4)
// 	pat2.PlaceJump(pat1)

// 	pat3.PlaceNoteClever(0x81, 16)
// 	pat3.PlaceNoteClever(0x82, 16)
// 	pat3.PlaceLoop(2, 2)
// 	pat3.PlaceLoop(1, 2)
// 	pat3.PlaceJump(pat2)

// 	song.SetInitialPattern(pat1, DAC)
// 	song.AttachPattern(pat2)
// 	song.AttachPattern(pat3)

// 	str := new(BinaryStringWriter)
// 	song.Export(str)
// 	str.Stop()

// 	fmt.Print(str.Release())
// }

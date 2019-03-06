package deflemask

// Instrument types
const (
	FM  = 1
	STD = 0
)

// Effect constants
const (
	FxArpeggio          = 0x00
	FxPortaUp           = 0x01
	FxPortaDown         = 0x02
	FxPortaNote         = 0x03
	FxVibrato           = 0x04
	FxVolSlidePortaNote = 0x05
	FxVolSlideVibrato   = 0x06
	FxTremolo           = 0x07
	FxPanning           = 0x08
	FxSetTickTime1      = 0x09
	FxVolSlide          = 0x0A
	FxPatternJump       = 0x0B
	FxRetrigger         = 0x0C
	FxPatternBreak      = 0x0D
	FxSetTickTime2      = 0x0F

	FxArpeggioLength = 0xE0
	FxPortaUpFix     = 0xE1
	FxPortaDownFix   = 0xE2
	FxVibratoMode    = 0xE3
	FxVibratoMaxAmp  = 0xE4
	FxFineTune       = 0xE5
	FxSetSampleBank  = 0xEB
	FxNoteCut        = 0xEC
	FxNoteDelay      = 0xED

	FxEnableDAC    = 0x17
	FxSetNoiseMode = 0x20
)

// Vibrato modes
const (
	VibNormal   = 0
	VibUpRect   = 1
	VibDownRect = 2
)

// Noise timbres
const (
	NTRandom   = 1
	NTPeriodic = 0
)

// Noise ranges
const (
	NRTritonic  = 0
	NRFullRange = 1
)

// System constants
const (
	SysGenesis            = 0x02
	SysGenesisFM3Ext      = 0x12
	SysMasterSystem       = 0x03
	SysGameBoy            = 0x04
	SysPCEngine           = 0x05
	SysNES                = 0x06
	SysCommodore64SID8580 = 0x07
	SysCommodore64SID6581 = 0x17
	SysArcadeYM2151       = 0x08
)

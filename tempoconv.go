package main

var tempoMap = map[int]int{
	30: 0x02,
	40: 0x03,
	45: 0x04,
	48: 0x05,
	50: 0x06,
	51: 0x07,
	52: 0x08,
	53: 0x09,
	54: 0x0A,
	55: 0x0C,
	56: 0x0F,
	57: 0x14,
	58: 0x1E,
	59: 0x3C,
	60: 0x00,
}

// GetOptimalTempo returns optimal tempo for specified FPS
func GetOptimalTempo(fps int) (tempoMod int) {
	lower := false
	equFPS := 0
	switch fps {
	case 31, 32, 33, 34:
		tempoMod = 2
		equFPS = 30

	case 35, 36, 37, 38, 39, 41, 42:
		tempoMod = 3
		equFPS = 40

	case 43, 44, 46:
		tempoMod = 4
		equFPS = 45

	case 47:
		tempoMod = 5
		equFPS = 48

	case 49:
		tempoMod = 6
		equFPS = 50

	default:
		if fps < 30 {
			tempoMod = 2
			equFPS = 30
			lower = true
		} else if fps > 60 {
			tempoMod = 0
			equFPS = 60
		} else {
			tempoMod = tempoMap[fps]
		}
	}

	if equFPS != 0 {
		logger.Printf(
			"warning: unable to find perfect tempo match for FPS = %d; will use tempo 01 %02X, which is equivalent for FPS = %d",
			fps,
			tempoMod,
			equFPS,
		)

		if lower {
			logger.Print(
				"hint: adjust FPS, base time and speed, so that FPS falls in range (30 : 60), but song plays at the same pace",
			)
		}
	}

	return
}

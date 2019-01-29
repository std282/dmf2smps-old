package freqs

import (
	"log"
	"os"
)

var logWarn = log.New(os.Stderr, "freqs: warning: ", 0)
var logError = log.New(os.Stderr, "freqs: error: ", 0)

package mngpat

import (
	"log"
	"os"
)

var logWarn = log.New(os.Stderr, "mngpat: warning: ", 0)

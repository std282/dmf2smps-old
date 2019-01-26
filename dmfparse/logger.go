package dmfparse

import (
	"log"
	"os"
)

var logWarn = log.New(os.Stderr, "dmfparse: warning: ", 0)
var logErr = log.New(os.Stderr, "dmfparse: error: ", 0)

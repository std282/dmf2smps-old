package deflemask

import (
	"io"
	"log"
	"os"
)

var logHint = log.New(os.Stderr, "deflemask: hint: ", 0)
var logWarn = log.New(os.Stderr, "deflemask: warning: ", 0)
var logErr = log.New(os.Stderr, "deflemask: error: ", 0)

// SetLogOutput set logging output to
func SetLogOutput(logOut io.Writer) {
	logHint.SetOutput(logOut)
	logWarn.SetOutput(logOut)
	logErr.SetOutput(logOut)
}

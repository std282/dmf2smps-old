package smpsbuild

/*
	This file contains loggers for errors and warnings, and function for log
	redirection.
*/

import (
	"io"
	"log"
	"os"
)

//var logger = log.New(os.Stdout, "smpsbuild: ", 0)

var logError = log.New(os.Stderr, "smpsbuild: error: ", 0)
var logWarn = log.New(os.Stdout, "smpsbuild: warning: ", 0)

// LogRedirect redirects logging to specified writer (stderr by default)
func LogRedirect(w io.Writer) {
	logError.SetOutput(w)
	logWarn.SetOutput(w)
}

package smpsbuild

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "smps: ", 0)

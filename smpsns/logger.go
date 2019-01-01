package smpsns

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "smpsns: ", 0)

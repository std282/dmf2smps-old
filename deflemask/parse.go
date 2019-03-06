package deflemask

import (
	"compress/zlib"
	"io"
	"os"
)

// Parse parses DMF file, filling it up
func Parse(dmfPath string) (dmf *Module, err error) {
	defer func() {
		err = internalRecover()
	}()

	dmfFile, err := os.OpenFile(dmfPath, os.O_RDONLY, 0)
	if err != nil {
		return
	}

	dmfFileDecoded, err := zlib.NewReader(dmfFile)
	if err != nil {
		return
	}

	pr := newPanicReader(dmfFileDecoded)
	dmf = new(Module)
	dmf.parse(pr)
	return
}

func (dmf *Module) parse(r io.Reader) {
	// TODO
}

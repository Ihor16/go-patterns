package strategy

import (
	"fmt"
	"os"
)

type tarCompressor struct {
}

func newTarCompressor() *tarCompressor {
	return &tarCompressor{}
}

func (z tarCompressor) Compress(file os.File) os.File {
	fmt.Println("compressing to tar")
	return file
}

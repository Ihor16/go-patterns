package strategy

import (
	"fmt"
	"os"
)

type zipCompressor struct {
}

func newZipCompressor() *zipCompressor {
	return &zipCompressor{}
}

func (z zipCompressor) Compress(file os.File) os.File {
	fmt.Println("compressing to zip")
	return file
}

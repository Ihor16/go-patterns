package strategy

import (
	"os"
)

type CompressionStrategy interface {
	Compress(file os.File) os.File
}

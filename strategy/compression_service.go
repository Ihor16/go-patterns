package strategy

import (
	"os"
)

type CompressionService struct {
	strategy CompressionStrategy
}

func NewCompressionService() *CompressionService {
	s := createStrategy()
	return &CompressionService{strategy: s}
}

func (c CompressionService) Compress(file os.File) os.File {
	return c.strategy.Compress(file)
}

package strategy

func createStrategy() CompressionStrategy {
	if isUnix() {
		return newTarCompressor()
	}
	return newZipCompressor()
}

func isUnix() bool {
	return true
}

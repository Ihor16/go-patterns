package strategy

type CompressionStrategyFactory struct {
}

func createStrategy() CompressionStrategy {
	if isUnix() {
		return newTarCompressor()
	}
	return newZipCompressor()
}

func isUnix() bool {
	return true
}

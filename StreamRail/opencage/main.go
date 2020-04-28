package main

import (
	"ConfluentADI/StreamRail/processing"
	"ConfluentADI/StreamRail/processor"
)

func main() {
	processor.RunProcessorTest(processing.Test) // press ctrl-c to stop
}

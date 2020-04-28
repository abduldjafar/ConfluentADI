package main

import (
	"ConfluentADI/StreamRail/processing"
	"ConfluentADI/StreamRail/processor"
)

func main() {

	//processor.RunProcessorOpenCage(processing.OpenCage) // press ctrl-c to stop
	processor.RunProcessorTiploc(processing.TiplocV1)
}

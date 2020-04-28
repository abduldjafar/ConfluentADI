package processor

import (
	"ConfluentADI/StreamRail/defineGroup"
	"ConfluentADI/config"
	"context"
	"github.com/lovoo/goka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func getconfig() *config.Configuration {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)
	return baseConfig
}

func process(proc *goka.Processor, err error) {
	p := proc
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	go func() {
		defer close(done)
		if err = p.Run(ctx); err != nil {
			log.Fatalf("error running processor: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait   // wait for SIGINT/SIGTERM
	cancel() // gracefully stop processor
	<-done
}

func RunProcessorOpenCage(streamProc func(ctx goka.Context, msg interface{})) {
	g := defineGroup.DefineGroupOpenCage(goka.Group(getconfig().Kafka.Gtopic), goka.Stream(getconfig().Kafka.Topic), streamProc)
	p, err := goka.NewProcessor([]string{getconfig().Kafka.Bootstrap}, g)
	process(p, err)

}

func RunProcessorTiploc(streamProc func(ctx goka.Context, msg interface{})) {
	g := defineGroup.DefineGroupTiploc(goka.Group(getconfig().Kafka.Gtopic), goka.Stream(getconfig().Kafka.Topic), streamProc)
	p, err := goka.NewProcessor([]string{getconfig().Kafka.Bootstrap}, g)
	process(p, err)
}

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

var (
	brokers             = []string{"localhost:9092"}
	topic   goka.Stream = "abduls"
	group   goka.Group  = "abduls-group"
)

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

func RunProcessorTest(streamProc func(ctx goka.Context, msg interface{})) {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	g := defineGroup.DefineGroupTest(goka.Group(baseConfig.Kafka.Gtopic), goka.Stream(baseConfig.Kafka.Topic), streamProc)
	p, err := goka.NewProcessor([]string{baseConfig.Kafka.Bootstrap}, g)
	process(p, err)

}

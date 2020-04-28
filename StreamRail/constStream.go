package StreamRail

import "github.com/lovoo/goka"

var (
	brokers             = []string{"localhost:9092"}
	topic   goka.Stream = "abduls"
	group   goka.Group  = "abduls-group"
)

func Vars() ([]string, goka.Stream, goka.Group) {
	return brokers, topic, group
}

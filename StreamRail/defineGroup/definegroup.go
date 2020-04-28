package defineGroup

import (
	"ConfluentADI/model/opencage"
	"ConfluentADI/model/tiploc"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

func definegroup(group goka.Group, topic goka.Stream, data func(ctx goka.Context, msg interface{}),
	input goka.Codec, output goka.Codec) *goka.GroupGraph {

	return goka.DefineGroup(group,
		goka.Input(topic, input, data),
		goka.Persist(output))
}

func DefineGroupOpenCage(group goka.Group, topic goka.Stream, data func(ctx goka.Context, msg interface{})) *goka.GroupGraph {
	return definegroup(group, topic, data, new(codec.String), new(opencage.ResultOpenCage))

}

func DefineGroupTiploc(group goka.Group, topic goka.Stream, data func(ctx goka.Context, msg interface{})) *goka.GroupGraph {
	return definegroup(group, topic, data, new(codec.String), new(tiploc.RequestTiploc))
}

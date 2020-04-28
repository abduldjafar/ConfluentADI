package defineGroup

import (
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

func DefineGroupTest(group goka.Group, topic goka.Stream, data func(ctx goka.Context, msg interface{})) *goka.GroupGraph {
	return goka.DefineGroup(group,
		goka.Input(topic, new(codec.String), data),
		goka.Persist(new(codec.Int64)),
	)
}

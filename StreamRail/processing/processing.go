package processing

import (
	"github.com/lovoo/goka"
	"log"
)

func Test(ctx goka.Context, msg interface{}) {
	var counter int64
	// ctx.Value() gets from the group table the value that is stored for
	// the message's key.
	if val := ctx.Value(); val != nil {
		counter = val.(int64)
	}
	counter++
	// SetValue stores the incremented counter in the group table for in
	// the message's key.
	ctx.SetValue(counter)
	log.Printf("key = %s, counter = %v, msg = %v", ctx.Key(), counter, msg)
}

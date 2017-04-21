package basic

import (
	"github.com/stretchr/testify/assert"
	"github.com/johnmcdnl/xrun/carrot"
	"fmt"
	"context"
)

func init() {

	carrot.Step(`^this scenario should execute 1 time and pass$`, func(ctx context.Context) {
		ctx = carrot.AddData(ctx, "userId", 10)
	})

	carrot.Step(`^setup was called 1 time$`, func(ctx context.Context) {
		fmt.Println(carrot.GetData(ctx, "userId"))
	})

	carrot.Step(`^I have an initial step$`, func(ctx context.Context) {
		assert.Equal(carrot.T(ctx), 4, 4)
	})

	carrot.Step(`^I have a second step$`, func(ctx context.Context) {
		assert.Fail(carrot.T(ctx), "Just Failing because why not")
	})

	carrot.Step(`^someone called "(.*)" jumps$`, func(ctx context.Context, user string) {
		assert.Fail(carrot.T(ctx), "Just Failing because why not")
	})

	carrot.Step(`^someone called "(.*)" jumps$`, func(ctx context.Context, user string) {
		fmt.Println("hello")
	})
}


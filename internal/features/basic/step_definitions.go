package basic

import (
	. "github.com/johnmcdnl/xrun/xrun"
	"github.com/stretchr/testify/assert"
	"github.com/johnmcdnl/xrun/carrot"
	"fmt"
)

func init() {


	carrot.Step(`^this scenario should execute 1 time and pass$`, func(tCtx *TestContext) {
		assert.Equal(tCtx.T(), 4, 4)
	})

	carrot.Step(`^setup was called 1 time$`, func(tCtx *TestContext) {
		assert.Equal(tCtx.T(), 4, 4)
	})

	carrot.Step(`^I have an initial step$`, func(tCtx *TestContext) {
		assert.Equal(tCtx.T(), 4, 4)
	})

	carrot.Step(`^I have a second step$`, func(tCtx *TestContext) {
		assert.Fail(tCtx.T(), "Just Failing because why not")
	})

	//someone called "Jack" jumps
	carrot.Step(`^someone called "(.*)" jumps$`, func(tCtx *TestContext, user string) {
		assert.Fail(tCtx.T(), "Just Failing because why not")
	})

	carrot.Step(`^someone called "(.*)" jumps$`, func(tCtx *TestContext, user string) {
		fmt.Println("hello")
	})
}

package basic

import (
	. "github.com/johnmcdnl/xrun/xrun"
	"github.com/stretchr/testify/assert"
)

func init() {
	And(`^I have an initial step$`, func(tCtx *TestContext) {
		assert.Equal(tCtx.T(), 4, 4)
	})

	And(`^I have a second step$`, func(tCtx *TestContext) {
		assert.Fail(tCtx.T(), "Just Failing because why not")
	})

	//someone called "Jack" jumps
	And(`^someone called "(.*)" jumps$`, func(tCtx *TestContext, user string) {
		assert.Fail(tCtx.T(), "Just Failing because why not")
	})
}

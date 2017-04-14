package basic

import (
	. "github.com/gucumber/gucumber"
	"github.com/stretchr/testify/assert"
)

func init() {
	executions := 100
	result := 1
	beforeAllCalls := 0

	BeforeAll(func() {
		beforeAllCalls++
	})

	Before("@basic", func() {
		executions = 0
	})

	Given(`^I have an initial step$`, func(tCtx TestContext) {

		assert.Equal(tCtx.T(), 1, 1)
	})

	And(`^I have a second step$`, func(tCtx TestContext) {

		assert.Equal(tCtx.T(), 2, 2)
	})

	When(`^I run the "(.+?)" command$`, func(tCtx TestContext, s1 string) {

		assert.Equal(tCtx.T(), "gucumber", s1)
	})

	Then(`^this scenario should execute (\d+) time and pass$`, func(tCtx TestContext, i1 int) {

		//executions++
		//assert.Equal(tCtx.T(), i1, executions)

	})

	Given(`^I perform (\d+) \+ (\d+)$`, func(tCtx TestContext, i1 int, i2 int) {

		result = i1 + i2
	})

	Then(`^I should get (\d+)$`, func(tCtx TestContext, i1 int) {

		assert.Equal(tCtx.T(), result, 4)
	})

	Then(`^setup was called (\d+) times?$`, func(tCtx TestContext, i2 int) {

		assert.Equal(tCtx.T(), i2, beforeAllCalls)


	})

	And(`^john made a wall of data$`, func(tCtx TestContext, data string) {

	})

	And(`^john has some data table$`, func(tCtx TestContext, table [][]string) {

	})

}

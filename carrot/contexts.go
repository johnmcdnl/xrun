package carrot

import (
	"context"
	"github.com/johnmcdnl/xrun/xrun"
)

func T(ctx context.Context) *xrun.TestingT {
	return ctx.Value("testingT").(*xrun.TestingT)
}

func AddData(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetData(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}

package filter


import (
	"context"
	"fmt"
)

type MyContext struct {
	context.Context
	KeyValue map[string]bool
}

type FilterFunc func(*MyContext) bool

type FilterFuncChain []FilterFunc

type CombinedFunc struct {
	CF    FilterFuncChain
	MyCtx *MyContext
}

func main() {
	myContext := MyContext{Context: context.TODO(), KeyValue: map[string]bool{"key": false}}

	cf := CombinedFilter(&myContext, F1, F2, F3);
	DoFilter(cf)
}

func DoFilter(cf *CombinedFunc) {
	for _, f := range cf.CF {
		res := f(cf.MyCtx)
		fmt.Println("result:", res)
		if res == false {
			fmt.Println("stopped")
			return
		}
	}
}

func CombinedFilter(ctx *MyContext, ff ...FilterFunc) *CombinedFunc {
	return &CombinedFunc{
		CF:    ff,
		MyCtx: ctx,
	}
}

func F1(ctx *MyContext) bool {
	ctx.KeyValue["key"] = true
	fmt.Println(ctx.KeyValue["key"])

	return ctx.KeyValue["key"]
}

func F2(ctx *MyContext) bool {
	ctx.KeyValue["key"] = false
	fmt.Println(ctx.KeyValue["key"])

	return ctx.KeyValue["key"]
}

func F3(ctx *MyContext) bool {
	ctx.KeyValue["key"] = false
	fmt.Println(ctx.KeyValue["key"])

	return ctx.KeyValue["key"]
}


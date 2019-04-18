package process_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/coc1961/process"
)

func TestProcess_OkTest(t *testing.T) {

	ctx := 0
	proc := process.New()

	proc.AddStep(Step1)
	proc.AddStep(Step2)
	proc.AddStep(Step3)

	proc.Start(&ctx)
	proc.RunAll()

	result := proc.Result()

	//assert.NoError(t, proc.Error())
	//assert.Equal(t, 3, result.(int))

}

/*
func TestProcess_TestWIthError(t *testing.T) {

	ctx := 0
	proc := process.New()

	proc.AddStep(Step1)
	proc.AddStep(BadStep)
	proc.AddStep(Step2)
	proc.AddStep(Step3)

	proc.Start(&ctx)
	proc.RunAll()

	result := proc.Result()
	assert.Error(t, proc.Error())
	assert.Equal(t, 1, result.(int))

}
*/

func Step1(ctx process.Context) (interface{}, error) {
	_ = ctx.(int)

	sCtx := ctx.(*int)
	*sCtx++
	fmt.Println("Step1 context", *sCtx)
	return *sCtx, nil
}
func Step2(ctx process.Context) (interface{}, error) {
	sCtx := ctx.(*int)
	*sCtx++
	fmt.Println("Step2 context", *sCtx)
	return *sCtx, nil
}
func Step3(ctx process.Context) (interface{}, error) {
	sCtx := ctx.(*int)
	*sCtx++
	fmt.Println("Step3 context", *sCtx)
	return *sCtx, nil
}
func BadStep(ctx process.Context) (interface{}, error) {
	sCtx := ctx.(*int)
	fmt.Println("BadStep context", *sCtx)
	return *sCtx, errors.New("Step Error")
}

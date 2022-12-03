// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Timer Operations

package gtimer_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhwei820/gcron/gtimer"
)

var ctx = context.Background()

func TestTimer_Add_Close(t *testing.T) {
	timer := gtimer.New()
	array := []int{}
	//fmt.Println("start", time.Now())
	timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
		//fmt.Println("job1", time.Now())
		array = append(array, 1)
	})
	timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
		//fmt.Println("job2", time.Now())
		array = append(array, 1)
	})
	timer.Add(ctx, 400*time.Millisecond, func(ctx context.Context) {
		//fmt.Println("job3", time.Now())
		array = append(array, 1)
	})
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 5)
	timer.Close()
	time.Sleep(250 * time.Millisecond)
	fixedLength := len(array)
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), fixedLength)
}

func TestTimer_Start_Stop_Close(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.Add(ctx, 1000*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	assert.Equal(t, len(array), 0)
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	timer.Stop()
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	timer.Start()
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	timer.Close()
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

func TestJob_Reset(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	job := timer.AddSingleton(ctx, 500*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(300 * time.Millisecond)
	job.Reset()
	time.Sleep(300 * time.Millisecond)
	job.Reset()
	time.Sleep(300 * time.Millisecond)
	job.Reset()
	time.Sleep(600 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestTimer_AddSingleton(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.AddSingleton(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
		time.Sleep(10 * time.Second)
	})
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 1)

	time.Sleep(500 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestTimer_AddOnce(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.AddOnce(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	timer.AddOnce(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	timer.Close()
	time.Sleep(250 * time.Millisecond)
	fixedLength := len(array)
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), fixedLength)
}

func TestTimer_AddTimes(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.AddTimes(ctx, 200*time.Millisecond, 2, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

func TestTimer_DelayAdd(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.DelayAdd(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestTimer_DelayAddJob(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.DelayAddEntry(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	}, false, 100, gtimer.StatusReady)
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestTimer_DelayAddSingleton(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.DelayAddSingleton(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
		time.Sleep(10 * time.Second)
	})
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 0)

	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestTimer_DelayAddOnce(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.DelayAddOnce(ctx, 200*time.Millisecond, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 0)

	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 1)

	time.Sleep(500 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestTimer_DelayAddTimes(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.DelayAddTimes(ctx, 200*time.Millisecond, 500*time.Millisecond, 2, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(200 * time.Millisecond)
	assert.Equal(t, len(array), 0)

	time.Sleep(600 * time.Millisecond)
	assert.Equal(t, len(array), 1)

	time.Sleep(600 * time.Millisecond)
	assert.Equal(t, len(array), 2)

	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

func TestTimer_AddLessThanInterval(t *testing.T) {

	timer := gtimer.New(gtimer.TimerOptions{
		Interval: 100 * time.Millisecond,
	})
	array := []int{}
	timer.Add(ctx, 20*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(50 * time.Millisecond)
	assert.Equal(t, len(array), 0)

	time.Sleep(110 * time.Millisecond)
	assert.Equal(t, len(array), 1)

	time.Sleep(110 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

func TestTimer_AddLeveledJob1(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.DelayAdd(ctx, 1000*time.Millisecond, 1000*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(1500 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	time.Sleep(1300 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestTimer_Exit(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
		gtimer.Exit()
	})
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

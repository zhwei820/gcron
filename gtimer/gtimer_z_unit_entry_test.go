// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Job Operations

package gtimer_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhwei820/gcron/gtimer"
)

func TestJob_Start_Stop_Close(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	job := timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	job.Stop()
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	job.Start()
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	job.Close()
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 2)

	assert.Equal(t, job.Status(), gtimer.StatusClosed)
}

func TestJob_Singleton(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	job := timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
		time.Sleep(10 * time.Second)
	})
	assert.Equal(t, job.IsSingleton(), false)
	job.SetSingleton(true)
	assert.Equal(t, job.IsSingleton(), true)
	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 1)

	time.Sleep(250 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestJob_SetTimes(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	job := timer.Add(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	job.SetTimes(2)
	//job.IsSingleton()
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

func TestJob_Run(t *testing.T) {

	timer := gtimer.New()
	array := []int{}
	job := timer.Add(ctx, 1000*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	job.Job()(ctx)
	assert.Equal(t, len(array), 1)
}

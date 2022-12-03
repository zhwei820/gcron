// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package functions

package gtimer_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhwei820/gcron/gtimer"
)

// var (
// 	ctx = context.TODO()
// )

func TestSetTimeout(t *testing.T) {

	array := []int{}
	gtimer.SetTimeout(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestSetInterval(t *testing.T) {

	array := []int{}
	gtimer.SetInterval(ctx, 300*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 3)
}

func TestAddEntry(t *testing.T) {

	array := []int{}
	gtimer.AddEntry(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	}, false, 2, gtimer.StatusReady)
	time.Sleep(1100 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

func TestAddSingleton(t *testing.T) {

	array := []int{}
	gtimer.AddSingleton(ctx, 200*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
		time.Sleep(10000 * time.Millisecond)
	})
	time.Sleep(1100 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestAddTimes(t *testing.T) {

	array := []int{}
	gtimer.AddTimes(ctx, 200*time.Millisecond, 2, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

func TestDelayAdd(t *testing.T) {

	array := []int{}
	gtimer.DelayAdd(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(600 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	time.Sleep(600 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestDelayAddEntry(t *testing.T) {

	array := []int{}
	gtimer.DelayAddEntry(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	}, false, 2, gtimer.StatusReady)
	time.Sleep(500 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	time.Sleep(2000 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

func TestDelayAddSingleton(t *testing.T) {

	array := []int{}
	gtimer.DelayAddSingleton(ctx, 500*time.Millisecond, 500*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
		time.Sleep(10000 * time.Millisecond)
	})
	time.Sleep(300 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestDelayAddOnce(t *testing.T) {

	array := []int{}
	gtimer.DelayAddOnce(ctx, 1000*time.Millisecond, 2000*time.Millisecond, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(2000 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	time.Sleep(2000 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestDelayAddTimes(t *testing.T) {

	array := []int{}
	gtimer.DelayAddTimes(ctx, 500*time.Millisecond, 500*time.Millisecond, 2, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(300 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	time.Sleep(1500 * time.Millisecond)
	assert.Equal(t, len(array), 2)
}

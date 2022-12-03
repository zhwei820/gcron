// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcron_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhwei820/gcron"
	"github.com/zhwei820/log"
)

var (
	ctx = context.TODO()
)

func init() {
	log.InitLogger("test", true, "debug", 3)
}
func TestCron_Add_Close(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	_, err1 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
		log.InfoZ(ctx, "cron1")
		array = append(array, 1)
	})
	_, err2 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
		log.InfoZ(ctx, "cron2")
		array = append(array, 1)
	}, "test")
	assert.Equal(t, err1, nil)
	assert.Equal(t, err2, nil)
	assert.Equal(t, cron.Size(), 2)
	time.Sleep(1300 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	time.Sleep(1300 * time.Millisecond)
	assert.Equal(t, len(array), 4)
	cron.Close()
	time.Sleep(1300 * time.Millisecond)
	fixedLength := len(array)
	time.Sleep(1300 * time.Millisecond)
	assert.Equal(t, len(array), fixedLength)
}

func TestCron_Basic(t *testing.T) {

	cron := gcron.New()
	cron.Add(ctx, "* * * * * *", func(ctx context.Context) {}, "add")
	// fmt.Println("start", time.Now())
	cron.DelayAdd(ctx, time.Second, "* * * * * *", func(ctx context.Context) {}, "delay_add")
	assert.Equal(t, cron.Size(), 1)
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, cron.Size(), 2)

	cron.Remove(gcron.GenCtx(), "delay_add")
	assert.Equal(t, cron.Size(), 1)

	entry1 := cron.Search("add")
	entry2 := cron.Search("test-none")
	assert.NotNil(t, entry1)
	assert.Nil(t, entry2)

	// test @ error

	cron = gcron.New()
	defer cron.Close()
	_, err := cron.Add(ctx, "@aaa", func(ctx context.Context) {}, "add")
	assert.NotNil(t, err)

	// test @every error

	cron = gcron.New()
	defer cron.Close()
	_, err = cron.Add(ctx, "@every xxx", func(ctx context.Context) {}, "add")
	assert.NotNil(t, err)

}

func TestCron_Remove(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
		array = append(array, 1)
	}, "add")
	assert.Equal(t, len(array), 0)
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, len(array), 1)

	cron.Remove(gcron.GenCtx(), "add")
	assert.Equal(t, len(array), 1)
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, len(array), 1)
}

func TestCron_Add_FixedPattern(t *testing.T) {
	for i := 0; i < 5; i++ {
		doTestCronAddFixedPattern(t)
	}
}

func doTestCronAddFixedPattern(t *testing.T) {

	var (
		now     = time.Now()
		cron    = gcron.New()
		array   = []int{}
		minutes = now.Minute()
		seconds = now.Second() + 2
	)
	defer cron.Close()

	if seconds >= 60 {
		seconds %= 60
		minutes++
	}
	var pattern = fmt.Sprintf(
		`%d %d %d %d %d %s`,
		seconds, minutes, now.Hour(), now.Day(), now.Month(), now.Weekday().String(),
	)
	log.DebugZ(ctx, fmt.Sprintf(`pattern: %s`, pattern))
	_, err := cron.Add(ctx, pattern, func(ctx context.Context) {
		array = append(array, 1)
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(3000 * time.Millisecond)
	log.DebugZ(ctx, `current time`)
	assert.Equal(t, len(array), 1)
}

func TestCron_AddSingleton(t *testing.T) {
	// un used, can be removed

	cron := gcron.New()
	cron.Add(ctx, "* * * * * *", func(ctx context.Context) {}, "add")
	cron.DelayAdd(ctx, time.Second, "* * * * * *", func(ctx context.Context) {}, "delay_add")
	assert.Equal(t, cron.Size(), 1)
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, cron.Size(), 2)

	cron.Remove(gcron.GenCtx(), "delay_add")
	assert.Equal(t, cron.Size(), 1)

	entry1 := cron.Search("add")
	entry2 := cron.Search("test-none")
	assert.NotNil(t, entry1)
	assert.Nil(t, entry2)
	// keep this

	cron = gcron.New()
	array := []int{}
	cron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
		array = append(array, 1)
		time.Sleep(50 * time.Second)
	})
	assert.Equal(t, cron.Size(), 1)
	time.Sleep(3500 * time.Millisecond)
	assert.Equal(t, len(array), 1)

}

func TestCron_AddOnce1(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	cron.AddOnce(ctx, "* * * * * *", func(ctx context.Context) {
		array = append(array, 1)
	})
	cron.AddOnce(ctx, "* * * * * *", func(ctx context.Context) {
		array = append(array, 1)
	})
	assert.Equal(t, cron.Size(), 2)
	time.Sleep(2500 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	assert.Equal(t, cron.Size(), 0)
}

func TestCron_AddOnce2(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	cron.AddOnce(ctx, "@every 2s", func(ctx context.Context) {
		array = append(array, 1)
	})
	assert.Equal(t, cron.Size(), 1)
	time.Sleep(3000 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	assert.Equal(t, cron.Size(), 0)
}

func TestCron_AddTimes(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	_, _ = cron.AddTimes(ctx, "* * * * * *", 2, func(ctx context.Context) {
		array = append(array, 1)
	})
	time.Sleep(3500 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	assert.Equal(t, cron.Size(), 0)
}

func TestCron_DelayAdd(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	cron.DelayAdd(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
		array = append(array, 1)
	})
	assert.Equal(t, cron.Size(), 0)
	time.Sleep(800 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	assert.Equal(t, cron.Size(), 1)
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	assert.Equal(t, cron.Size(), 1)
}

func TestCron_DelayAddSingleton(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	cron.DelayAddSingleton(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
		array = append(array, 1)
		time.Sleep(10 * time.Second)
	})
	assert.Equal(t, cron.Size(), 0)
	time.Sleep(2200 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	assert.Equal(t, cron.Size(), 1)
}

func TestCron_DelayAddOnce(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	cron.DelayAddOnce(ctx, 500*time.Millisecond, "* * * * * *", func(ctx context.Context) {
		array = append(array, 1)
	})
	assert.Equal(t, cron.Size(), 0)
	time.Sleep(800 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	assert.Equal(t, cron.Size(), 1)
	time.Sleep(2200 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	assert.Equal(t, cron.Size(), 0)
}

func TestCron_DelayAddTimes(t *testing.T) {

	cron := gcron.New()
	array := []int{}
	cron.DelayAddTimes(ctx, 500*time.Millisecond, "* * * * * *", 2, func(ctx context.Context) {
		array = append(array, 1)
	})
	assert.Equal(t, cron.Size(), 0)
	time.Sleep(800 * time.Millisecond)
	assert.Equal(t, len(array), 0)
	assert.Equal(t, cron.Size(), 1)
	time.Sleep(3000 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	assert.Equal(t, cron.Size(), 0)
}

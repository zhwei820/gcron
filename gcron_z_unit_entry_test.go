// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcron_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhwei820/gcron"
	"github.com/zhwei820/log"
)

func TestCron_Entry_Operations(t *testing.T) {
	log.InitLogger("test", true, "debug", 3)

	var (
		cron  = gcron.New()
		array = []int{}
	)
	ctx := context.Background()
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

	cron = gcron.New()
	array = []int{}
	entry, err1 := cron.Add(ctx, "* * * * * *", func(ctx context.Context) {
		array = append(array, 1)
	})
	assert.Equal(t, err1, nil)
	assert.Equal(t, len(array), 0)
	assert.Equal(t, cron.Size(), 1)
	time.Sleep(1300 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	assert.Equal(t, cron.Size(), 1)
	entry.Stop()
	time.Sleep(5000 * time.Millisecond)
	assert.Equal(t, len(array), 1)
	assert.Equal(t, cron.Size(), 1)
	entry.Start()
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, len(array), 2)
	assert.Equal(t, cron.Size(), 1)
	entry.Close()
	time.Sleep(1200 * time.Millisecond)
	assert.Equal(t, cron.Size(), 0)
}

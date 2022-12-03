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

	"github.com/zhwei820/gcron"
	"github.com/zhwei820/log"
)

func Test_cronAdd(t *testing.T) {
	log.InitLogger("test", true, "debug", 3)

	// c, err := gcron.NewEtcdMutexBuilder(etcdclient.Config{
	// 	Endpoints: []string{"127.0.0.1:2379"},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	cron := gcron.New()

	cron.Add(ctx, "*/3 * * * * *", func(ctx context.Context) {
		log.InfoZ(context.TODO(), "doing")
		// time.Sleep(500 * time.Millisecond)
	})
	time.Sleep(10 * time.Second)
}

func Test_example(t *testing.T) {
	cron := gcron.New()
	array := []int64{}
	cron.AddTimes(ctx, "@every 2s", 1, func(ctx context.Context) {
		array = append(array, 1)
	}, "cron1")
	cron.AddOnce(ctx, "@every 2s", func(ctx context.Context) {
		array = append(array, 1)
	}, "cron2")
	fmt.Println(len(array), cron.Size())
	cron.Remove(gcron.GenCtx(), "cron2")
	fmt.Println(len(array), cron.Size())
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(len(array), cron.Size())
	// Output:
	// 0 2
	// 0 1
	// 1 0
}

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

	"github.com/zhwei820/gcron/gcron"
	"github.com/zhwei820/log"
)

func Test_cronAddSingleton(t *testing.T) {
	log.InitLogger("test", true, "debug", 3)
	gcron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
		log.InfoZ(context.TODO(), "doing")
		time.Sleep(500 * time.Millisecond)
	})
	select {}
}

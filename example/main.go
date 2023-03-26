package main

import (
	"context"
	"time"

	"github.com/zhwei820/gcron"
	"github.com/zhwei820/log"
)

func main() {
	log.InitLogger("test", true, log.RunModeDebug, 3)
	ctx := gcron.GenCtx()
	log.InfoZ(ctx, "start....")

	cron := gcron.NewWithETCD("127.0.0.1:2379", "biz2")

	_, err := cron.Add(ctx, "*/3 * * * * *", func(ctx context.Context) {
		log.InfoZ(ctx, "doing")
		time.Sleep(500 * time.Millisecond)
	}, "demo_task_id")
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	// cron.Stop(ctx, "demo_task_id")
	// cron.Remove(ctx, "demo_task_id")

	_, err = cron.Add(ctx, "*/3 * * * * *", func(ctx context.Context) {
		log.InfoZ(ctx, "doing")
		time.Sleep(500 * time.Millisecond)
	}, "demo_task_id2")
	if err != nil {
		panic(err)
	}
	time.Sleep(300 * time.Second)

}

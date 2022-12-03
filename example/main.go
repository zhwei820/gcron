package main

import (
	"context"
	"time"

	etcdclient "go.etcd.io/etcd/client/v3"

	"github.com/zhwei820/gcron"
	"github.com/zhwei820/log"
)

func main() {
	log.InitLogger("test", true, "debug", 3)
	ctx := gcron.GenCtx()

	c, err := gcron.NewEtcdMutexBuilder(&etcdclient.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	cron := gcron.New(gcron.WithEtcdMutexBuilder(c))

	_, err = cron.Add(ctx, "*/3 * * * * *", func(ctx context.Context) {
		log.InfoZ(ctx, "doing")
		time.Sleep(500 * time.Millisecond)
	}, "demo_task_id")
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	cron.Stop(ctx, "demo_task_id")
	cron.Remove(ctx, "demo_task_id")

	time.Sleep(30 * time.Second)

}

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
	ctx := context.Background()

	c, err := gcron.NewEtcdMutexBuilder(&etcdclient.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	cron := gcron.New(gcron.WithEtcdMutexBuilder(c))

	_, err = cron.Add(ctx, "*/3 * * * * *", func(ctx context.Context) {
		log.InfoZ(context.TODO(), "doing")
		time.Sleep(500 * time.Millisecond)
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(60 * time.Second)
}

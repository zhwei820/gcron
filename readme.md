# gcron

A distributed and fault tolerant cron runner in golang.

1. fork from github.com/gogf/gf/v2/os/gcron
2. with etcd election support (ref: github.com/zhwei820/election)

## Goal

This package aims at implementing a distributed and fault tolerant cron in order to:

- Run an identical process on several hosts
- Each of these process instantiate a cron with the same rules
- Ensure only one of these processes executes an iteration of a job

## Example

```go
    log.InitLogger("test", true, "debug", 3)
	ctx := context.Background()
	cron := gcron.NewWithETCD("127.0.0.1:2379", "biz2") // use etcd election, guarantee that only one cron process runs at a time.

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

```

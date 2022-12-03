# gcron

A distributed and fault tolerant cron runner in golang.

1. fork from github.com/gogf/gf/v2/os/gcron
2. with etcd support (ref:  github.com/Scalingo/go-etcd-cron)


## Goal

This package aims at implementing a distributed and fault tolerant cron in order to:

* Run an identical process on several hosts
* Each of these process instantiate a cron with the same rules
* Ensure only one of these processes executes an iteration of a job

## Example

```go
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

```
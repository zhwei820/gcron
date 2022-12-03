package gcron

import (
	"context"

	"github.com/google/uuid"
	"github.com/zhwei820/log"
)

func GenCtx() context.Context {
	return context.WithValue(context.Background(), log.TraceID, uuid.NewString())
}

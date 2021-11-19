package pkg

import (
	"context"
	"log"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Hello(ctx context.Context) {
	span, ctx := tracer.StartSpanFromContext(ctx, "hello")
	defer span.Finish()
	time.Sleep(time.Second * 5)
	log.Print("hello")
}

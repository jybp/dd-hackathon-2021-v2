package main

import (
	"context"

	"github.com/jybp/dd-hackathon-2021/pkg"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	rules := []tracer.SamplingRule{tracer.RateRule(1)}
	tracer.Start(
		tracer.WithSamplingRules(rules),
		tracer.WithService("jb-hackathon-2021"),
		tracer.WithEnv("dev"),
	)
	defer tracer.Stop()
	ctx := context.Background()
	for {
		span, ctx := tracer.StartSpanFromContext(ctx, "main")
		pkg.Hello(ctx)
		span.Finish()
	}
}

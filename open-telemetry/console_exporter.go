package open_telemetry

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	_ "go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func ConsoleExporter () {

	otelgin
	exporter, err := stdout.NewExporter(stdout.WithPrettyPrint(), )
	if err != nil {
		log.Fatalf("failed to initialize stdout exporter pipeline: %v", err)
	}
	ctx := context.Background()
	defer exporter.Shutdown(ctx)
	// Creating a Tracer Provider
	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))
	// Handle this error in a sensible manner where possible
	defer func() { _ = tp.Shutdown(ctx) }()

	// Creating a Meter Provider
	pusher := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			exporter,
		),
		controller.WithPusher(exporter),
		controller.WithCollectPeriod(5*time.Second),
	)

	err = pusher.Start(ctx)
	if err != nil {
		log.Fatalf("failed to initialize metric controller: %v", err)
	}

	// Handle this error in a sensible manner where possible
	defer func() { _ = pusher.Stop(ctx) }()

	otel.SetTracerProvider(tp)
	// otel.SetMeterProvider(pusher.MeterProvider())
	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(propagator)

	_, span := otel.GetTracerProvider().Tracer("client").Start(ctx, "spanName")
	defer span.End()
	attr1 := attribute.KeyValue{Key: attribute.Key("stringAttr"), Value: attribute.StringValue("Yes")}
	span.SetAttributes(attribute.String("stringAttr", "Yes"), attr1)
	span.AddEvent("Start do work")
	// do work
	span.AddEvent("work finished!!")

}

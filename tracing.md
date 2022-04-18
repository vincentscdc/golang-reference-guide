# Tracing

We can leverage opentelemetry, which is the new and upcoming merge of all previous tracing norms.
It should be backward compatible with opencensus and opentracing propagators.

## Intro

[ref](https://opentelemetry.io/)

## Lib

Use the official [go opentelemetry libs](https://github.com/open-telemetry/opentelemetry-go)

## Collector to send the spans to

You can send the spans you create to:

* New relic or sumologic directly
* A [local collector](https://github.com/open-telemetry/opentelemetry-collector)
* Any other service

## Calling other APIs

Spans and traces need to be propagated along all other services.

(TODO: opentelemetry http client)

## Middleware to add tracing your http service

This will need to be managed internally.

[Reference lib working with go-chi router](https://github.com/riandyrn/otelchi)

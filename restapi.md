# REST API

To keep things simple, we won't use GRPC and GRPC gateway for now.
As we want to keep the compatibility with other services in the company, it's best to stick (for now) to the simplest implementation of REST APIs.

## Naming things (MUST)

Follow REST principles: use names for entities, **NO VERB**, the verb is the HTTP method.

## Server (MUST)

Use the net/http one, but don't forget to set the different timeouts (read, idle...)
Also think about the graceful shutdown (requests inflight)

## Router (MUST)

Many choices here, from gorilla/mux to gin...
**My preference goes to [go-chi](https://go-chi.io/#/)**

One important point to note: try not to use a "context based" router, it does not play well with a lot of middlewares.

Use the major (semantic) version in the URL to allow a deterministic response type while the API can evolve.

😈: look fasthttprouter is on top of the benchmarks!
    sure... but it does not take stdlib http handlers, so it will prevent you from migrating to other router once yours is deprecated or not fashionable anymore.

## Middlewares

As long as your router respects the stdlib net/http way of doing things, you have a vast amount of choices.
Any middleware built around net/http will work.

It is highly recommended to use a middleware for tracing.

## Handlers

Handlers should be compatible with net/http stdlib, usually after being wrapped.

Do not use global vars to access other parts of the app, like the database.
Use dependency injection and build the handler as a returned closure directly when you define it.

### Encoding

Right now, we ll most probably only use JSON, but that should be done via the wrapper functions

### Wrapping (OPINION)

A small wrapping lib will be provided in order to enforce the error display and handling, as well as limit the boilerplate code.
(TODO: create this lib and link it here).

### Errors

Leverage the ErrorResponse to build your error response! it makes errors much clearer for the users of your APIs.
Also refer to the [errors chapter](./errors.md)

### Versioning (MUST)

We use semantic versioning.
Always use the major version in the URI of your endpoints, so that we can have multiple versions running together.
The versioning in the URL should follow the git tags and your application versioning.

### Swagger OPENAPI (MUST)

You will have to generate docs for each of your endpoints.
You can easily do that by leveraging [golang http-swagger](https://github.com/swaggo/http-swagger), compatible with the net/http routers (including go-chi).

## Need some SQL?

follow the [rdbms section](./rdbms.md)

## Need some cache or nosql?

follow the [nosql section](./nosql.md)

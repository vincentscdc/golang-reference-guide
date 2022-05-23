# non RDBMS

## Redis

### Packages to use

We use [go-redis](https://github.com/go-redis) library, which

- provides type-safe API for each redis command
- supports single server, sentinel, and cluster modes

go-redis uses pool management for each connections implicitly, so there is no need to manage the connection manually in most cases. Also it is rare to close the client while the application is living.

## Others (OPINION)

You probably don't need it, leverage postgres JSON, if you REALLY, REALLY can't know the data structure otherwise.

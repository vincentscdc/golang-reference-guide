# 0

```bash
    go run hi.go leaky.go main.go 
```

```bash
    curl -i 127.0.0.1:8081/hi
```

## pprof

```bash
curl "http://localhost:8081/debug/pprof/trace?seconds=5" > trace_nonleakyoptim.out
```

```bash
    k6 run --vus 100 --duration 2s bench_nonleakyoptim.js
```

```bash
go tool pprof -http=localhost:9090 "localhost:8081/debug/pprof/heap"
```

```bash
curl "http://localhost:8081/debug/pprof/trace?seconds=5" > trace_nonleaky.out
```

```bash
    k6 run --vus 100 --duration 2s bench_nonleaky.js
```

```bash
go tool pprof -http=localhost:9090 "localhost:8081/debug/pprof/heap"
```

```bash
curl "http://localhost:8081/debug/pprof/trace?seconds=5" > trace_leaky.out
```

```bash
    k6 run --vus 100 --duration 2s bench_leaky.js
```

```bash
go tool pprof -http=localhost:9090 "localhost:8081/debug/pprof/heap"
```

```bash
go tool trace -http=localhost:9090 trace_nonleakyoptim.out
go tool trace -http=localhost:9090 trace_nonleaky.out
go tool trace -http=localhost:9090 trace_leaky.out
```

```bash
go tool pprof -http=localhost:9090 "localhost:8081/debug/pprof/profile?seconds=8"
```

```bash
go tool pprof -http=localhost:9090 "localhost:8081/debug/pprof/goroutine?seconds=8"
```

```bash
go tool pprof -http=localhost:9090 "localhost:8081/debug/pprof/mutex?seconds=8"
```

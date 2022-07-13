# 0

```bash
    curl -i 127.0.0.1:8081/hi
```

## stats

go to [stats viewer](http://localhost:18066/debug/statsview)

```bash
    k6 run --vus 100 --duration 2s bench_nonleakyoptim.js
```

```bash
    k6 run --vus 100 --duration 2s bench_nonleaky.js
```

```bash
    k6 run --vus 100 --duration 2s bench_leaky.js
```

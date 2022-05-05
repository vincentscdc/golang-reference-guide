# RDBMS

## Drivers and connection

[lib/pq](https://github.com/lib/pq) is now in maintenance mode (see README).
Hence, use [pgx](https://github.com/jackc/pgx)

## Modeling, Querying (DDL, DML)

### The repository pattern

Please refer to this great article from [threedots.tech](https://threedots.tech/post/repository-pattern-in-go/).

In a few words: abstract your DB interactions from the business logic

### Packages to use

#### Selection criteria

* Development speed
* Cockroachdb compatibility
* Postgres compatibility
* Performance
* Type safety

#### Choices

* Type safe, fast, especially in simple CRUD dev: [sqlc](https://github.com/kyleconroy/sqlc)
caveat: for any case with joins and/or dynamic queries, do not use sqlc yet, it's not available [yet](https://github.com/kyleconroy/sqlc/discussions/363)

* Type safe, complex queries, lots of joins and dynamic queries (a better ORM): [sqlboiler](https://github.com/volatiletech/sqlboiler)
* Also a good choice, but a bit more boilerplate: [sqlx](https://github.com/jmoiron/sqlx)
* For the best perf, use raw queries with [pgx](https://github.com/jackc/pgx)

Learn SQL!

#### 😈 why not gorm / gorp / pop ... ?

Benches on the latest versions of differents ORMs:

```bash
goos: linux
goarch: amd64
pkg: github.com/volatiletech/boilbench [github.com]
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz

BenchmarkBoilDelete/boil-8       1980972               605.6 ns/op           152 B/op          7 allocs/op
BenchmarkGORPDelete/gorp-8       1153930               985.0 ns/op           336 B/op         12 allocs/op
BenchmarkGORMDelete/gorm-8        193442              6231 ns/op            4586 B/op         58 allocs/op
BenchmarkPOPDelete/pop-8          173594              7013 ns/op             760 B/op         72 allocs/op



BenchmarkBoilInsert/boil-8        687402              1750 ns/op             920 B/op         19 allocs/op
BenchmarkGORPInsert/gorp-8        411339              2876 ns/op            1368 B/op         31 allocs/op
BenchmarkGORMInsert/gorm-8         91862             12191 ns/op            7455 B/op         88 allocs/op
BenchmarkPOPInsert/pop-8           32780             35185 ns/op            8243 B/op        298 allocs/op



BenchmarkBoilRawBind/boil-8       200460              6010 ns/op            3938 B/op         34 allocs/op
BenchmarkGORPRawBind/gorp-8        57151             19899 ns/op            7219 B/op        218 allocs/op
BenchmarkPopRawBind/pop-8          92578             11745 ns/op            5127 B/op         50 allocs/op
BenchmarkGORMRawBind/gorm-8        70704             15539 ns/op            9681 B/op        103 allocs/op



BenchmarkBoilSelectAll/boil-8     154538              6718 ns/op            2869 B/op         47 allocs/op
BenchmarkGORMSelectAll/gorm-8      74418             15159 ns/op           10199 B/op         94 allocs/op
BenchmarkGORPSelectAll/gorp-8      57510             19792 ns/op            7219 B/op        218 allocs/op
BenchmarkPopSelectAll/pop-8        44373             26178 ns/op            6626 B/op        133 allocs/op



BenchmarkBoilSelectSubset/boil-8                  146150              7133 ns/op            3014 B/op         51 allocs/op
BenchmarkGORMSelectSubset/gorm-8                   71102             15761 ns/op           10320 B/op         96 allocs/op
BenchmarkGORPSelectSubset/gorp-8                   56043             20185 ns/op            7219 B/op        218 allocs/op
BenchmarkPopSelectSubset/pop-8                     47995             24042 ns/op            6268 B/op        136 allocs/op



BenchmarkBoilSelectComplex/boil-8                 134761              8847 ns/op            3871 B/op         71 allocs/op
BenchmarkGORMSelectComplex/gorm-8                  61298             18763 ns/op           11884 B/op        131 allocs/op
BenchmarkGORPSelectComplex/gorp-8                  54746             20036 ns/op            7508 B/op        221 allocs/op
BenchmarkPopSelectComplex/pop-8                    41752             27368 ns/op            7059 B/op        153 allocs/op



BenchmarkBoilUpdate/boil-8                        852603              1401 ns/op             880 B/op         15 allocs/op
BenchmarkGORPUpdate/gorp-8                        362690              2834 ns/op            1480 B/op         32 allocs/op
BenchmarkGORMUpdate/gorm-8                        126267              8235 ns/op            5955 B/op         64 allocs/op
BenchmarkPopUpdate/pop-8                           38670             30268 ns/op            7779 B/op        289 allocs/op
```

I think this speaks better than anything.

## Migrations

The reference is [golang-migrate](https://github.com/golang-migrate/migrate).

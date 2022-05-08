# Money

There is no money/decimal library in std golang, [at least for go 1.x](https://github.com/golang/go/issues/19787).

As of today, there are 4 full-featured arbitrary-precision decimal-floating point libraries for Go out there:
[Eric Lagergren's decimal](https://github.com/ericlagergren/decimal), [CockroachDB's apd](https://github.com/cockroachdb/apd), [Shopspring's decimal](https://github.com/shopspring/decimal) and [db47h's decimal](https://github.com/db47h/decimal).

Quoting [some existing benchmarks](https://github.com/db47h/decimal#performance):

| digits | 9 | 19 | 38 | 100 | 500 | 5000 |
|--------|--:|---:|---:|----:|----:|-----:|
| Eric's decimal (Go) | 6415 | 30254 | 65171 | 194263 | 1731528 | 89841923 |
| decimal | 12887 | 42720 | 100878 | 348865 | 4212811 | 342349031|
| Eric's decimal (GDA) | 7124 | 39357 | 107720 | 392453 | 5421146 | 1175936547 |
| Shopspring's decimal | 39528 | 96261 | 204017 | 561321 | 3402562 | 97370022 |
| apd | 70833 | 301098 | 1262021 | 9859180 | 716558666 | ??? |

## The constraints

### Storage

When you store your data in a RDMS, most likely postgres, in our case, you will translate your go decimal type to the NUMERIC/DECIMAL (same thing) in postgres.

### JSON

The type needs to be serialized/deserialized in a good fashion.

## What did other libraries used?

It turns out that sqlboiler chose [Eric Lagergren's decimal](https://github.com/ericlagergren/decimal) for its NUMERIC mapping.
The performance is great and it seems Eric Lagergren is actually an active participant in the future std lib implementation.

The other contender is [Shopspring's decimal](https://github.com/shopspring/decimal), which has been around for a while and which is widely used.

## Final decision

We will use [Eric Lagergren's decimal](https://github.com/ericlagergren/decimal) to represent arbitrary-precision decimal-floating point types as:

* It refers to the IEEE and ISO std implementation
* It is fast (anecdotal, as it will probably never be a bottleneck for us)
* The author is actively participating in the std lib implementation
* sqlboiler chose this library to map numeric

# go-kv-benchmark

Fork of [akrylysov/pogreb-bench](https://github.com/akrylysov/pogreb-bench)

Supports 

- [prologic/bitcask](https://git.mills.io/prologic/bitcask)
- [akrylysov/pogreb](https://github.com/akrylysov/pogreb)
- [arriqaaq/flashdb](https://github.com/arriqaaq/flashdb)
- [dgraph-io/badger/v3](https://github.com/dgraph-io/badger)
- [millken/archivedb](https://github.com/millken/archivedb)
- [millken/godb](https://github.com/millken/godb)
- [syndtr/goleveldb](https://github.com/syndtr/goleveldb)
- [tidwall/buntdb](https://github.com/tidwall/buntdb)
- [bbolt](https://go.etcd.io/bbolt)

## Usage:

```#!bash
$ sh run.sh
```

## Benchmarks (2025-04-08)

Benchmark results on a MacMini M1

| Engine     | Concurrency | Number of keys | Key Size | Value Size | Write/s  | Reads/s | Size      | Time (sec) |
|------------|-------------|----------------|----------|------------|----------|---------|-----------|------------|
| archivedb  | 10          | 100000         | 16~64    | 128~512    | 916929   | 3048397 | 1.00GB    | 0.142      |
| godb       | 1           | 100000         | 16~64    | 128~512    | 806124   | 1000046 | 40.00MB   | 0.224      |
| archivedb  | 1           | 100000         | 16~64    | 128~512    | 754325   | 1018488 | 1.00GB    | 0.231      |
| godb       | 10          | 100000         | 16~64    | 128~512    | 677299   | 1909193 | 76.00MB   | 0.200      |
| buntdb     | 1           | 100000         | 16~64    | 128~512    | 238888   | 826306  | 37.04MB   | 0.540      |
| flashdb    | 1           | 100000         | 16~64    | 128~512    | 201245   | 833126  | 36.61MB   | 0.617      |
| goleveldb  | 1           | 100000         | 16~64    | 128~512    | 298213   | 267387  | 35.62MB   | 0.709      |
| badger     | 10          | 100000         | 16~64    | 128~512    | 197553   | 556338  | 2.15GB    | 0.686      |
| badger     | 1           | 100000         | 16~64    | 128~512    | 129431   | 345569  | 2.14GB    | 1.062      |
| bitcask    | 1           | 100000         | 16~64    | 128~512    | 128437   | 610090  | 42.73MB   | 0.942      |
| pogreb     | 1           | 100000         | 16~64    | 128~512    | 143600   | 533163  | 37.83MB   | 0.884      |
| flashdb    | 10          | 100000         | 16~64    | 128~512    | 160399   | 3269875 | 73.21MB   | 0.654      |
| buntdb     | 10          | 100000         | 16~64    | 128~512    | 198750   | 2051855 | 74.08MB   | 0.552      |
| goleveldb  | 10          | 100000         | 16~64    | 128~512    | 269065   | 498853  | 71.35MB   | 0.572      |
| bitcask    | 10          | 100000         | 16~64    | 128~512    | 71869    | 899225  | 85.54MB   | 1.503      |
| pogreb     | 10          | 100000         | 16~64    | 128~512    | 106199   | 568768  | 75.60MB   | 1.117      |
| bbolt      | 1           | 100000         | 16~64    | 128~512    | 35284    | 705959  | 64.05MB   | 2.976      |
| bbolt      | 10          | 100000         | 16~64    | 128~512    | 30279    | 619775  | 112.09MB  | 3.464      |
| lmdb       | 1           | 100000         | 16~64    | 128~512    | 12143    | 454437  | 52.93MB   | 8.455      |
| lmdb       | 10          | 100000         | 16~64    | 128~512    | 9987     | 338420  | 105.71MB  | 10.308     |



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

| Engine    | Concurrency | Number of keys | Key Size   | Value Size | Write/s  | Reads/s  | Size (MB) | Time (sec) |
|-----------|-------------|----------------|------------|------------|----------|----------|-----------|------------|
| pogreb    | 1           | 100000         | 16~64      | 128~512    | 127434   | 528534   | 37.80     | 0.974      |
| goleveldb | 1           | 100000         | 16~64      | 128~512    | 298758   | 290917   | 35.57     | 0.678      |
| bbolt     | 1           | 100000         | 16~64      | 128~512    | 33137    | 864519   | 64.05     | 3.133      |
| badger    | 1           | 100000         | 16~64      | 128~512    | 145236   | 250687   | 2.14      | 1.087      |
| bitcask   | 1           | 100000         | 16~64      | 128~512    | 129872   | 734167   | 42.74     | 0.906      |
| archivedb | 1           | 100000         | 16~64      | 128~512    | 1811245  | 1214254  | 1.00      | 0.138      |
| flashdb   | 1           | 100000         | 16~64      | 128~512    | 221612   | 830516   | 36.62     | 0.572      |
| buntdb    | 1           | 100000         | 16~64      | 128~512    | 274587   | 1064432  | 37.04     | 0.458      |
| godb      | 1           | 100000         | 16~64      | 128~512    | 1049367  | 1006178  | 64.00     | 0.195      |
| lmdb      | 1           | 100000         | 16~64      | 128~512    | 12557    | 453581   | 52.69     | 8.184      |
| pogreb    | 10          | 100000         | 16~64      | 128~512    | 113766   | 554729   | 75.67     | 1.059      |
| goleveldb | 10          | 100000         | 16~64      | 128~512    | 205725   | 401808   | 71.26     | 0.735      |
| bbolt     | 10          | 100000         | 16~64      | 128~512    | 29490    | 802135   | 112.09    | 3.516      |
| badger    | 10          | 100000         | 16~64      | 128~512    | 216440   | 284171   | 2.15      | 0.814      |
| bitcask   | 10          | 100000         | 16~64      | 128~512    | 74828    | 854236   | 85.51     | 1.453      |
| archivedb | 10          | 100000         | 16~64      | 128~512    | 1016055  | 2755282  | 1.00      | 0.135      |
| flashdb   | 10          | 100000         | 16~64      | 128~512    | 169539   | 2940664  | 73.28     | 0.624      |
| buntdb    | 10          | 100000         | 16~64      | 128~512    | 210643   | 1952032  | 74.06     | 0.526      |
| godb      | 10          | 100000         | 16~64      | 128~512    | 968540   | 3225290  | 96.00     | 0.134      |
| lmdb      | 10          | 100000         | 16~64      | 128~512    | 10652    | 350766   | 105.24    | 9.673      |



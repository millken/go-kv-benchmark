# go-kv-benchmark

Fork of [akrylysov/pogreb-bench](https://github.com/akrylysov/pogreb-bench)

Supports 

- [prologic/bitcask](https://git.mills.io/prologic/bitcask)
- [akrylysov/pogreb](https://github.com/akrylysov/pogreb)
- [arriqaaq/flashdb](https://github.com/arriqaaq/flashdb)
- [dgraph-io/badger/v3](https://github.com/dgraph-io/badger)
- [millken/archivedb](https://github.com/millken/archivedb)
- [syndtr/goleveldb](https://github.com/syndtr/goleveldb)
- [tidwall/buntdb](https://github.com/tidwall/buntdb)
- [bbolt](https://go.etcd.io/bbolt)

## Usage:

```#!bash
$ sh run.sh
```

## Benchmarks (2021-03-01)

Benchmark results on a Macbook 11" Dual core Intel Core i7 with 16GB RAM (~2015 model).

|Engine   |Concurrency|Number of keys|Key Size|Value Size|Write/s|Reads/s|Size (MB)|Time (sec)|
|---------|-----------|--------------|--------|----------|-------|-------|---------|----------|
|pogreb   |1          |100000        |16~64   |128~512   |10118  |173209 |37.79MB  |10.460    |
|goleveldb|1          |100000        |16~64   |128~512   |20852  |102263 |35.65MB  |5.773     |
|bbolt    |1          |100000        |16~64   |128~512   |2933   |326489 |56.01MB  |34.394    |
|badger   |1          |100000        |16~64   |128~512   |19492  |154955 |2.14GB   |5.775     |
|bitcask  |1          |100000        |16~64   |128~512   |6205   |244413 |42.69MB  |16.525    |
|archivedb|1          |100000        |16~64   |128~512   |207241 |1948314|1.00GB   |0.534     |
|flashdb  |1          |100000        |16~64   |128~512   |20272  |559364 |36.57MB  |5.111     |
|buntdb   |1          |100000        |16~64   |128~512   |17736  |457576 |37.00MB  |5.857     |
|pogreb   |10         |100000        |16~64   |128~512   |8602   |928049 |75.67MB  |11.732    |
|goleveldb|10         |100000        |16~64   |128~512   |44996  |439251 |71.30MB  |2.450     |
|bbolt    |10         |100000        |16~64   |128~512   |2482   |545551 |119.36MB |40.468    |
|badger   |10         |100000        |16~64   |128~512   |62479  |437967 |2.15GB   |1.829     |
|bitcask  |10         |100000        |16~64   |128~512   |3622   |1009171|85.43MB  |27.701    |
|archivedb|10         |100000        |16~64   |128~512   |186069 |7911332|1.00GB   |0.550     |
|flashdb  |10         |100000        |16~64   |128~512   |15502  |3185850|73.18MB  |6.482     |
|buntdb   |10         |100000        |16~64   |128~512   |20918  |2035365|73.98MB  |4.830     |


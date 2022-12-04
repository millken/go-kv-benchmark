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
|pogreb   |1          |100000        |16~64   |128~512   |9189   |163367 |37.82MB  |11.494    |
|goleveldb|1          |100000        |16~64   |128~512   |20116  |95945  |35.65MB  |6.013     |
|bbolt    |1          |100000        |16~64   |128~512   |2770   |269796 |56.01MB  |36.464    |
|badger   |1          |100000        |16~64   |128~512   |30732  |150567 |2.14GB   |3.918     |
|bitcask  |1          |100000        |16~64   |128~512   |7000   |252127 |42.69MB  |14.681    |
|archivedb|1          |100000        |16~64   |128~512   |42013  |1233604|1.00GB   |2.461     |
|flashdb  |1          |100000        |16~64   |128~512   |21791  |706699 |36.66MB  |4.731     |
|buntdb   |1          |100000        |16~64   |128~512   |17961  |447034 |36.99MB  |5.791     |
|nutsdb   |1          |100000        |16~64   |128~512   |92704  |124632 |256.00MB |1.881     |
|pogreb   |10         |100000        |16~64   |128~512   |8460   |742462 |75.61MB  |11.954    |
|goleveldb|10         |100000        |16~64   |128~512   |47099  |387700 |71.28MB  |2.381     |
|bbolt    |10         |100000        |16~64   |128~512   |2512   |400781 |119.16MB |40.044    |
|badger   |10         |100000        |16~64   |128~512   |57714  |440929 |2.15GB   |1.959     |
|bitcask  |10         |100000        |16~64   |128~512   |3686   |718327 |85.43MB  |27.264    |
|archivedb|10         |100000        |16~64   |128~512   |221822 |8585402|1.00GB   |0.462     |
|flashdb  |10         |100000        |16~64   |128~512   |16768  |3445436|73.27MB  |5.993     |
|buntdb   |10         |100000        |16~64   |128~512   |16984  |2408618|73.99MB  |5.929     |
|nutsdb   |10         |100000        |16~64   |128~512   |67976  |1009535|256.00MB |1.570     |



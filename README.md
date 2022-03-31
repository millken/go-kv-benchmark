# bitcask-bench

Fork of [akrylysov/pogreb-bench](https://github.com/akrylysov/pogreb-bench)
with support for Go11Modules + [prologic/bitcask](https://git.mills.io/prologic/bitcask)

## Usage:

```#!bash
$ go get git.mills.io/prologic/bitcask-bench
$ bitcask-bench -d ./tmp -e bitcask
```

## Benchmarks

Benchmark results on a Macbook 11" Dual core Intel Core i7 with 16GB RAM (~2015 model).

| Engine        | Concurrency   | Reads/s  | Write/s |
| ------------- |:-------------:| --------:| -------:|
| pogreb        |      1        | 1028574  | 22838   |
| goleveldb     |      1        | 78180    | 78607   |
| bbolt         |      1        | 340504   | 13836   |
| badger        |      1        | 118662   | 28599   |
| bitcask       |      1        | 441101   | 14686   |
|               |               |          |         |
| pogreb        |      10       | 2011091  | 12805   |
| goleveldb     |      10       | 194571   | 82431   |
| bbolt         |      10       | 545263   | 7885    |
| badger        |      10       | 316895   | 69107   |
| bitcask       |      10       | 922374   | 11926   |

#!/bin/sh

ENGINES="pogreb goleveldb bbolt badger bitcask archivedb flashdb buntdb nutsdb godb"
GOROUTINES="1 10"

go build

echo "Engine,Concurrency,Number of keys, Key Size, Value Size,Write/s,Reads/s,Size (MB),Time (sec)" > results.csv
for goroutines in ${GOROUTINES}; do
  for engine in ${ENGINES}; do
    echo "Benchmarking ${engine} with ${goroutines} threads ..."
    ./go-kv-benchmark -d ./tmp -c "${goroutines}" -e "${engine}"
  done
done

rm -rf ./tmp
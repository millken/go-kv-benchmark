#!/bin/sh

ENGINES="pogreb goleveldb bbolt badger bitcask archivedb"
GOROUTINES="1 10"

go build

for goroutines in ${GOROUTINES}; do
  for engine in ${ENGINES}; do
    echo "Benchmarking ${engine} with ${goroutines} threads ..."
    ./go-kv-benchmark -d ./tmp -c "${goroutines}" -e "${engine}"
  done
done

rm -rf ./tmp
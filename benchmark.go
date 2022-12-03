package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

func randKey(minL int, maxL int) string {
	n := rand.Intn(maxL-minL+1) + minL
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = byte(rand.Intn(25) + 65)
	}
	return string(buf)
}

func randValue(rnd *rand.Rand, src []byte, minS int, maxS int) []byte {
	n := rnd.Intn(maxS-minS+1) + minS
	return src[:n]
}

func forceGC() {
	runtime.GC()
	time.Sleep(time.Millisecond * 500)
}

func shuffle(a [][]byte) {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func generateKeys(count int, minL int, maxL int) [][]byte {
	keys := make([][]byte, 0, count)
	seen := make(map[string]struct{}, count)
	for len(keys) < count {
		k := randKey(minL, maxL)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		keys = append(keys, []byte(k))
	}
	return keys
}

func concurrentBatch(keys [][]byte, concurrency int, cb func(gid int, batch [][]byte) error) error {
	wg := &sync.WaitGroup{}
	batchSize := len(keys) / concurrency
	wg.Add(concurrency)
	var err error
	for i := 0; i < concurrency; i++ {
		batchStart := i * batchSize
		batchEnd := (i + 1) * batchSize
		if batchEnd > len(keys) {
			batchEnd = len(keys)
		}
		go func(gid int, batch [][]byte) {
			err = cb(gid, batch)
			wg.Done()
		}(i, keys[batchStart:batchEnd])
	}
	wg.Wait()
	return err
}

func showProgress(gid int, i int, total int) {
	if i%50000 == 0 {
		fmt.Printf("Goroutine %d. Processed %d from %d items...\n", gid, i, total)
	}
}

func benchmark(engine string, dir string, numKeys int, minKS int, maxKS int, minVS int, maxVS int, concurrency int, progress bool) error {
	ctr, err := getEngineCtr(engine)
	if err != nil {
		return err
	}

	dbpath := path.Join(dir, "bench_"+engine)
	db, err := ctr(dbpath)
	if err != nil {
		return err
	}

	fmt.Printf("Number of keys: %d\n", numKeys)
	fmt.Printf("Minimum key size: %d, maximum key size: %d\n", minKS, maxKS)
	fmt.Printf("Minimum value size: %d, maximum value size: %d\n", minVS, maxVS)
	fmt.Printf("Concurrency: %d\n", concurrency)
	fmt.Printf("Running %s benchmark...\n", engine)

	keys := generateKeys(numKeys, minKS, maxKS)
	valSrc := make([]byte, maxVS)
	if _, err := rand.Read(valSrc); err != nil {
		return err
	}
	forceGC()

	start := time.Now()
	err = concurrentBatch(keys, concurrency, func(gid int, batch [][]byte) error {
		rnd := rand.New(rand.NewSource(int64(rand.Uint64())))
		for i, k := range batch {
			if err := db.Put(k, randValue(rnd, valSrc, minVS, maxVS)); err != nil {
				return err
			}
			if progress {
				showProgress(gid, i, len(batch))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	endsecs := time.Since(start).Seconds()
	totalalsecs := endsecs
	writeNums := int(float64(numKeys) / endsecs)
	fmt.Printf("Put: %.3f sec, %d ops/sec\n", endsecs, writeNums)

	shuffle(keys)
	db, err = ctr(dbpath)
	if err != nil {
		return err
	}
	forceGC()

	start = time.Now()
	err = concurrentBatch(keys, concurrency, func(gid int, batch [][]byte) error {
		for i, k := range batch {
			v, err := db.Get(k)
			if err != nil {
				return err
			}
			if v == nil {
				return errors.New("key doesn't exist")
			}
			if progress {
				showProgress(gid, i, len(batch))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	endsecs = time.Since(start).Seconds()
	totalalsecs += endsecs
	readNums := int(float64(numKeys) / endsecs)
	fmt.Printf("Get: %.3f sec, %d ops/sec\n", endsecs, readNums)
	fmt.Printf("Put + Get time: %.3f sec\n", totalalsecs)
	sz, err := db.FileSize()
	if err != nil {
		return err
	}
	fmt.Printf("File size: %s\n", byteSize(sz))

	f, err := os.OpenFile("results.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	csvResult := fmt.Sprintf("%s,%d,%d,%d~%d,%d~%d,%d,%d,%s,%.3f\n", engine, concurrency, numKeys, minKS, maxKS, minVS, maxVS, writeNums, readNums, byteSize(sz), totalalsecs)
	if _, err := f.WriteString(csvResult); err != nil {
		return err
	}

	return db.Close()
}

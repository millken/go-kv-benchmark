package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/pkg/errors"
	"github.com/pkg/profile"
)

var (
	engine       = flag.String("e", "bitcask", "database engine name. pogreb, goleveldb, bbolt, badger, or bitcask")
	numKeys      = flag.Int("n", 100000, "number of keys")
	minKeySize   = flag.Int("mink", 16, "minimum key size")
	maxKeySize   = flag.Int("maxk", 64, "maximum key size")
	minValueSize = flag.Int("minv", 128, "minimum value size")
	maxValueSize = flag.Int("maxv", 512, "maximum value size")
	concurrency  = flag.Int("c", 1, "number of concurrent goroutines")
	dir          = flag.String("d", "", "database directory")
	progress     = flag.Bool("p", false, "show progress")
	profileMode  = flag.String("profile", "", "enable profile. cpu, mem, block or mutex")
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	var startMemStats, endMemStats runtime.MemStats
	runtime.ReadMemStats(&startMemStats)

	flag.Parse()

	if *dir == "" {
		flag.Usage()
		return
	}

	err := os.MkdirAll(*dir, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch *profileMode {
	case "cpu":
		defer profile.Start(profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile).Stop()
	case "mutex":
		defer profile.Start(profile.MutexProfile).Stop()
	}

	if err := benchmark(*engine, *dir, *numKeys, *minKeySize, *maxKeySize, *minValueSize, *maxValueSize, *concurrency, *progress); err != nil {
		fmt.Fprintf(os.Stderr, "\n%+v\n", errors.WithStack(err))
	}
	runtime.ReadMemStats(&endMemStats)
	fmt.Printf("MemoryStats: Alloc = %v MiB", bToMb(endMemStats.Alloc-startMemStats.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(endMemStats.TotalAlloc-startMemStats.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(endMemStats.Sys-startMemStats.Sys))
	fmt.Printf("\tNumGC = %v\n", endMemStats.NumGC-startMemStats.NumGC)

}

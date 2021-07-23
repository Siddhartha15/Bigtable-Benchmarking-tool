package loadtestbigtable

import (
	"benchmark/stat"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"amagi.com/dam/commons/logger"
)

var (
	log = logger.NewLogger("component", "benchmark_client")
)

// 	csvOutput 		- output path for statistics in .csv format. If this file already exists it will be overwritten
//  populate		- BatchWrite if true
// 	runFor			- how long to run the load test for; 0 to run forever until SIGTERM
// 	reqCount 		- number of concurrent requests
//  populateCount 	- number if rows to insert in bigtable

func LoadtestReads(csvOutput string, populate bool, runFor time.Duration,
	reqCount int, populateCount int64, keyRange int64) {

	var err error
	if populate {
		log.Info("START bulk writing")
		startTime := time.Now()
		BatchWrite(populateCount)
		endTime := time.Now()
		log.Infof("END bulk writing: Done within %f seconds", endTime.Sub(startTime).Seconds())
	}

	var csvFile *os.File
	if csvOutput != "" {
		csvFile, err = os.Create(csvOutput)
		if err != nil {
			log.Fatal("creating csv output file: %v", err)
		}
		defer csvFile.Close()
		log.Infof("Writing statistics to %q ...", csvOutput)
	}

	log.Infof("Starting load test... (run for %v)", runFor)
	sem := make(chan int, reqCount) // limit the number of requests happening at once
	var reads stats
	startTime := time.Now()
	stopTime := startTime.Add(runFor)
	var wg sync.WaitGroup
	for time.Now().Before(stopTime) || runFor == 0 {
		sem <- 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			rowKey := generateRandomRowkey(keyRange) // get random key within [0,keyRange)
			// rowKey := "dummy"

			ok := true
			opStart := time.Now()
			var stats *stats
			defer func() {
				stats.Record(ok, time.Since(opStart))
			}()

			stats = &reads
			ok = ReadSingleRow(rowKey)

		}()
	}
	wg.Wait()

	readsAgg := stat.NewAggregate("reads", reads.ds, reads.tries-reads.ok)
	log.Infof("Reads (%d ok / %d tries):%v", reads.ok, reads.tries, readsAgg)
	endTime := time.Now()
	log.Infof("Reads time taken: %f", endTime.Sub(startTime).Seconds())

	if csvFile != nil {
		stat.WriteCSV([]*stat.Aggregate{readsAgg}, csvFile)
	}
}

var allStats int64 // atomic

type stats struct {
	mu        sync.Mutex
	tries, ok int
	ds        []time.Duration
}

func (s *stats) Record(ok bool, d time.Duration) {
	s.mu.Lock()
	s.tries++
	if ok {
		s.ok++
	}
	s.ds = append(s.ds, d)
	s.mu.Unlock()

	if n := atomic.AddInt64(&allStats, 1); n%1000 == 0 {
		log.Infof("Progress: done %d ops", n)
	}
}

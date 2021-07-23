package main

import (
	loadtestbigtable "benchmark/loadtestBigtable"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"amagi.com/dam/commons/conffetch"
	"github.com/spf13/viper"
)

func main() {
	file, _ := ioutil.ReadFile("./testData/config.toml")
	reader := strings.NewReader(string(file))
	viper.SetConfigType(conffetch.ConfigType)
	_ = viper.ReadConfig(reader)

	emulate := flag.Bool("emulate", false, "set bigtable env if true")
	runFor := flag.Int("run_for", 1, "how long to run the load test for; 0 to run forever until SIGTERM")
	populateCount := flag.Int64("populate_count", 10, "populate table with populateCount")
	csvOutput := flag.String("csv_output", "reads_default.csv", "output path for statistics in .csv format. If this file already exists it will be overwritten.")
	reqCount := flag.Int("req_count", 1, "number of concurrent requests")
	populate := flag.Bool("populate", false, "populate bigtable if true")
	keyRange := flag.Int64("key_range", 10, "generate rowkey within [0, keyRange)")
	flag.Parse()

	if *emulate {
		os.Setenv("BIGTABLE_EMULATOR_HOST", "localhost:8086") // BigTable emulator env init
	}
	if *populate {
		*csvOutput = fmt.Sprintf("reads_t_%d_%ds_%d_%d.csv", *reqCount, *runFor, *populateCount, *keyRange)
	} else {
		*csvOutput = fmt.Sprintf("reads_f_%d_%ds_%d.csv", *reqCount, *runFor, *keyRange)
	}
	filePath, _ := filepath.Abs("./loadtestOutput/" + *csvOutput)

	loadtestbigtable.LoadtestReads(filePath, *populate, time.Duration(*runFor)*time.Second,
		*reqCount, *populateCount, *keyRange)
}

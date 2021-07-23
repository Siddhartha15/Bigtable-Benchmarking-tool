# Bigtable-benchmarking

## Install
  `$ go mod vendor`

  `$ go mod tidy`
 
## Build
  `$ go build -o bin/`
  
## Benchmark
  `$ ./benchmark -populate=true -populate_count=1000000 -req_count=3000 -key_range=2000000 -run_for=900`
    
## Info
* populate - populate bigtable if `-populate=true`
*	populate_count - populate bigtable with populateCount 
    * Ex: `-populate_count=1000000`, populates for 1Million rows
* req_count - number of concurrent requests 
  * Ex: `-req_count=3000`, with 3000 concurrent requests
* key_range - generate rowkey within `[0, keyRange)`
*	run_for - how long to run the load test for; 0 to run forever until SIGTERM
  * Ex: `-run_for=900`, run for 900 seconds

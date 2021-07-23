# Bigtable-benchmarking-tool

## Install
  `$ go mod vendor`

  `$ go mod tidy`
 
## Build
  `$ go build -o bin/`
  
## Benchmark
  `$ ./benchmark -populate=true -populate_count=1000000 -req_count=3000 -key_range=2000000 -run_for=900`
    
## Info
* __populate__ - populate bigtable if `-populate=true`

*	__populate_count__ - populate bigtable with populateCount 
    * Ex: `-populate_count=1000000`, populates for 1Million rows
   
* __req_count__ - number of concurrent requests 
  * Ex: `-req_count=3000`, with 3000 concurrent requests
  
* __key_range__ - generate rowkey within `[0, keyRange)`

*	__run_for__ - how long to run the load test for; 0 to run forever until SIGTERM
    * Ex: `-run_for=900`, run for 900 seconds

### For longer loadtest
* Run inside cloud instance, as background process

# Bigtable-benchmarking-tool

## Build
  `$ go mod vendor`

  `$ go mod tidy`
 
  `$ go build -o bin/`
  
## Run
  `$ ./benchmark -populate=true -populate_count=1000000 -req_count=3000 -key_range=2000000 -run_for=900`
 
## Results 
* [Bigtable-Performance](https://github.com/Siddhartha15/bigtable-benchmarking-tool/blob/main/loadtestOutput/bigtable_Performance.csv)
    
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

module benchmark

replace amagi.com/dam/commons/logger => ../../commons/logger

replace amagi.com/dam/commons/conffetch => ../../commons/conffetch

go 1.16

require (
	amagi.com/dam/commons/conffetch v0.0.0-00010101000000-000000000000
	amagi.com/dam/commons/logger v0.0.0-00010101000000-000000000000
	cloud.google.com/go/bigtable v1.10.1
	github.com/spf13/viper v1.8.1
	google.golang.org/api v0.47.0
)

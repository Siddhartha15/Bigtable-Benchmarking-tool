package loadtestbigtable

import (
	"benchmark/clients"
	"benchmark/config"
	"context"

	"cloud.google.com/go/bigtable"
)

func ReadRows(rowKeys []string) error {
	tableName := config.GetTranscodeTable()
	columnFamilyNames := config.GetColumnFamilyNames()

	ctx := context.Background()
	client := clients.GetBigTableClient()
	tbl := client.Open(tableName)

	columnName := columnFamilyNames[0] // transcoded_ads_cf -- refer config.toml

	// log.Infof("Reading all rows:")
	err := tbl.ReadRows(ctx, bigtable.PrefixRange(columnName), func(row bigtable.Row) bool {
		item := row[columnFamilyNames[0]][0]
		log.Info("Read Rows", "rowkey:", item.Row, "row value:", string(item.Value))
		return true
	}, bigtable.RowFilter(bigtable.ColumnFilter(columnName)))

	if err != nil {
		log.Fatal(err.Error(), "msg", "Could not read all rows ")
		return err
	}
	return nil
}

func ReadSingleRow(rowKey string) bool {

	ctx := context.Background()
	client := clients.GetBigTableClient()
	tbl := client.Open(config.GetTranscodeTable())

	// log.Info("Getting a single row ")
	row, err := tbl.ReadRow(ctx, rowKey)
	if err != nil {
		log.Info(err.Error(), "msg", "Could not read row with key ")
		return false
	}
	if len(row) == 0 {
		log.Info("Row not present")
		return false
	}

	// log.Info("Read Row", "rowkey:", rowKey, "row value:", string(row["transcoded_ads_cf"][1].Value))
	return true
}

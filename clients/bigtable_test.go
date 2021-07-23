package clients

import (
	"context"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/bigtable"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func ExampleGetBigTableClient() {
	// add row
	client := GetBigTableClient()
	tableClient := client.Open("test_table")
	addition := bigtable.NewMutation()
	addition.Set("transcoded_ads_cf", "key", bigtable.Now(), []byte("profile_id#version"))
	addition.Set("transcoded_ads_cf", "url", bigtable.Now(), []byte("adurl"))
	if err := tableClient.Apply(context.Background(), "row_key", addition); err != nil {
		fmt.Printf("Error writing to big table . %v", err.Error())
	}

	// read row
	row, err := tableClient.ReadRow(context.Background(), "row_key", bigtable.RowFilter(bigtable.ColumnFilter("transcoded_ads_cf")))
	if err != nil {
		panic(err)
	}
	for key, val := range row {
		fmt.Println("Column family" + key)
		for _, item := range val {
			fmt.Println(string(item.Value)) // value
		}
	}
}

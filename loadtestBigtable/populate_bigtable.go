package loadtestbigtable

import (
	"benchmark/clients"
	"benchmark/config"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"cloud.google.com/go/bigtable"
)

const (
	adUrl  = "https://www.tempurl.com/temp"
	key    = "profile1#v0.0.1"
	status = "RECIEVED"
)

func BatchWrite(populateCount int64) {

	tableName := config.GetTranscodeTable()

	ctx := context.Background()
	client := clients.GetBigTableClient()
	tbl := client.Open(tableName)

	var muts []*bigtable.Mutation

	for i := int64(0); i < populateCount; i++ {

		timestamp := bigtable.Now()
		mut := bigtable.NewMutation()
		mut.Set("transcoded_ads_cf", "key", timestamp, []byte(key))
		mut.Set("transcoded_ads_cf", "url", timestamp, []byte(adUrl))
		mut.Set("transcoded_status_cf", "key", timestamp, []byte(key))
		mut.Set("transcoded_status_cf", "status", timestamp, []byte(status))
		muts = append(muts, mut)

	}

	rowKeys := generateArrayOfRowkeys(populateCount)

	if errs, err := tbl.ApplyBulk(ctx, rowKeys, muts); err != nil {
		log.Fatal("Error bullk applying all rows", "error:", err)
	} else if errs != nil {
		log.Fatal("Error bullk applying some rows", "error:", errs)
	}

}
func SimpleWrite(populateCount int64) {

	tableName := config.GetTranscodeTable()

	ctx := context.Background()
	client := clients.GetBigTableClient()
	tbl := client.Open(tableName)

	log.Infof("SART simple writing")
	startTime := time.Now()
	rowKeys := generateArrayOfRowkeys(populateCount)

	for i := int64(0); i < populateCount; i++ {

		timestamp := bigtable.Now()
		mut := bigtable.NewMutation()
		mut.Set("transcoded_ads_cf", "key", timestamp, []byte(key))
		mut.Set("transcoded_ads_cf", "url", timestamp, []byte(adUrl))
		mut.Set("transcoded_status_cf", "key", timestamp, []byte(key))
		mut.Set("transcoded_status_cf", "status", timestamp, []byte(status))
		if err := tbl.Apply(ctx, rowKeys[i], mut); err != nil {
			log.Fatal("Error simple writing this row", "error:", err, "rowkey", rowKeys[i])

		}
	}
	endTime := time.Now()
	log.Infof("END simple writing: Done within %f seconds", endTime.Sub(startTime).Seconds())

}

//generates array of rowkeys from [0,populateCount)
func generateArrayOfRowkeys(populateCount int64) []string {

	rowKeys := make([]string, populateCount)

	for i := int64(0); i < populateCount; i++ {
		seedStr := strconv.FormatUint(uint64(i), 10)
		h := sha256.New()
		h.Write([]byte(seedStr))
		sha := h.Sum(nil)
		rowKeys[i] = hex.EncodeToString(sha)
	}

	return rowKeys
}

// generate random rowkey from [0,populateCount)
func generateRandomRowkey(populateCount int64) string {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	in := r1.Int63n(int64(populateCount))
	// log.Info("Random Rowkey", "Int Vlaue:", in)
	seedStr := strconv.FormatUint(uint64(in), 10)
	h := sha256.New()
	h.Write([]byte(seedStr))
	sha := h.Sum(nil)
	rowkey := hex.EncodeToString(sha) // String representation
	// log.Info("rand hash", "out", shaStr)
	return rowkey
}

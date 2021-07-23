package loadtestbigtable

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"amagi.com/dam/commons/conffetch"
	"github.com/spf13/viper"
)

const ()

func TestMain(m *testing.M) {
	file, err := ioutil.ReadFile("../testData/config.toml")
	if err != nil {
		log.Fatal(err.Error(), "Error reading config file")
	}
	viper.Set("test_environment", true)
	reader := strings.NewReader(string(file))
	viper.SetConfigType(conffetch.ConfigType)
	_ = viper.ReadConfig(reader)
	os.Setenv("BIGTABLE_EMULATOR_HOST", "localhost:8086") // BigTable env init
	os.Exit(m.Run())
}

func TestBatchWrite(t *testing.T) {
	populateCount := int64(10) // populate table with populateCount
	BatchWrite(populateCount)
}

func TestSimpleWrite(t *testing.T) {
	populateCount := int64(10) // populate table with populateCount
	SimpleWrite(populateCount)
}

func TestReadSingleRow(t *testing.T) {
	rowKey := "fffff7f18e3f2477c5c981222df6260c01b0e9324cfc8758f3c1cb9e9a920d79" // dummy key
	// rowKey := generateRandomRowkey(1000000)
	ReadSingleRow(rowKey)
}

func TestReadAllRows(t *testing.T) {
	rowKeys := []string{"dummy1", "dummy2"} //dummy values
	ReadRows(rowKeys)
}

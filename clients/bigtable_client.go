package clients

import (
	"context"
	"sync"

	"benchmark/config"

	"amagi.com/dam/commons/logger"
	"cloud.google.com/go/bigtable"
	"google.golang.org/api/option"
)

var (
	btOnce = &sync.Once{}
	client *bigtable.Client
	log    = logger.NewLogger("component", "bigtable_client")
)

// sliceContains reports whether the provided string is present in the given slice of strings.
func sliceContains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}

func createTableIfNotExists() {
	ctx := context.Background()
	log.Info("msg", "instance:", config.GetBigTableInstance(), "projectid:", config.GetProjectID())
	adminClient, err := bigtable.NewAdminClient(ctx, config.GetProjectID(), config.GetBigTableInstance())
	if err != nil {
		log.Fatal(err.Error(), "msg", "error creating client")
		panic(err)
	}

	// [START bigtable_hw_create_table]
	tables, err := adminClient.Tables(ctx)
	if err != nil {
		log.Fatal(err.Error(), "msg", "error listing tables")
	}

	if !sliceContains(tables, config.GetTranscodeTable()) {
		log.Infof("Creating table %s", config.GetTranscodeTable())
		if err := adminClient.CreateTable(ctx, config.GetTranscodeTable()); err != nil {
			log.Fatal(err.Error(), "msg", "Could not create table", "tableName:", config.GetTranscodeTable())
			panic(err)
		}
	}

	tblInfo, err := adminClient.TableInfo(ctx, config.GetTranscodeTable())
	if err != nil {
		log.Fatal(err.Error(), "msg", "Could not read info for table", "tableName:", config.GetTranscodeTable())
	}

	for _, columnFamilyName := range config.GetColumnFamilyNames() {
		if !sliceContains(tblInfo.Families, columnFamilyName) {
			if err := adminClient.CreateColumnFamily(ctx, config.GetTranscodeTable(), columnFamilyName); err != nil {
				log.Fatal(err.Error(), "msg", "Could not create column family", "Column Family:", columnFamilyName)
			}
		}
	}
	// [END bigtable_hw_create_table]

}

func GetBigTableClient() *bigtable.Client {
	btOnce.Do(func() {
		createTableIfNotExists()
		var err error
		var options []option.ClientOption

		log.Info("grpc connection pool", "pool size", config.GetGrpcConnPoolSize())
		options = append(options,
			option.WithGRPCConnectionPool(config.GetGrpcConnPoolSize()),

			// TODO(grpc/grpc-go#1388) using connection pool without WithBlock
			// can cause RPCs to fail randomly. We can delete this after the issue is fixed.
			// option.WithGRPCDialOption(grpc.WithBlock())
		)

		client, err = bigtable.NewClient(context.Background(), config.GetProjectID(), config.GetBigTableInstance(), options...)
		if err != nil {
			log.Fatal(err.Error(), "msg", "error getting client")
			panic(err)
		}
	})
	return client
}

func DeleteTableIfExists() {
	ctx := context.Background()
	adminClient, err := bigtable.NewAdminClient(ctx, config.GetProjectID(), config.GetBigTableInstance())
	if err != nil {
		log.Fatal(err.Error(), "msg", "error creating client")
		panic(err)
	}

	tables, err := adminClient.Tables(ctx)
	if err != nil {
		log.Fatal(err.Error(), "msg", "error listing tables")
		panic(err)
	}
	found := false
	for _, table := range tables {
		if table == config.GetTranscodeTable() {
			found = true
			break
		}
	}
	if found {
		log.Warn("required table present", "Deleting the table", "table", config.GetTranscodeTable())
		err := adminClient.DeleteTable(ctx, config.GetTranscodeTable())
		if err != nil {
			log.Fatal(err.Error(), "msg", "error deleting table", "table", config.GetTranscodeTable())
			panic(err)
		}
	}

	if err = adminClient.Close(); err != nil {
		log.Fatal(err.Error(), "msg", "error closing client")
		panic(err)
	}
	// [END bigtable_hw_delete_table]
}

package config

import (
	"github.com/spf13/viper"
)

func GetProjectID() string {
	return viper.GetString("project_id")
}

func GetBigTableInstance() string {
	return viper.GetString("big_table_instance")
}

func GetTranscodeTable() string {
	return viper.GetString("transcode_table")
}

func GetColumnFamilyNames() []string {
	return viper.GetStringSlice("column_family_names")
}

func GetColumnNamesForEachFamily(columnFamilyName string) []string {
	return viper.GetStringSlice(columnFamilyName)
}

func GetTranscodeJobTopic() string {
	return viper.GetString("transcoder_job_topic")
}

func GetTranscodeJobSubscription() string {
	return viper.GetString("transcoder_job_subscription")
}
func GetGrpcConnPoolSize() int {
	return viper.GetInt("conn_poolsize")
}

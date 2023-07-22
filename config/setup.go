package config

import (
	"github.com/spf13/viper"
	"os"
)

func Setup() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")

	viper.SetDefault("token", os.Getenv("CRONIQ_TOKEN"))

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	_ = viper.UnmarshalKey("token", &Token)
	_ = viper.UnmarshalKey("cronjobs", &CronJobs)
	_ = viper.UnmarshalKey("queues", &Queues)
}

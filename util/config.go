package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	CrawlWorkerSize   int    `mapstructure:"CRAWL_WORKER_SIZE"`
	CrawlBlockChunk   int64  `mapstructure:"CRAWL_BLOCK_CHUNK"`
	CrawlStartNum     int    `mapstructure:"CRAWL_START_NUM"`
	MigrationURL      string `mapstructure:"MIGRATION_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

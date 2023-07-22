package config

var (
	Token    string
	CronJobs []cronjob
	Queues   []queue
)

type cronjob struct {
	Name     string                 `mapstructure:"name"`
	Pattern  string                 `mapstructure:"pattern"`
	Endpoint string                 `mapstructure:"endpoint"`
	Data     map[string]interface{} `mapstructure:"data"`
}

type queue struct {
	Name     string `mapstructure:"name"`
	Endpoint string `mapstructure:"endpoint"`
}

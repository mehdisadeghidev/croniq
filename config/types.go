package config

var (
	Token    string
	CronJobs []cronjob
)

type cronjob struct {
	Name     string                 `mapstructure:"name"`
	Pattern  string                 `mapstructure:"pattern"`
	Endpoint string                 `mapstructure:"endpoint"`
	Data     map[string]interface{} `mapstructure:"data"`
}

package cron

import (
	"github.com/go-resty/resty/v2"
	"github.com/mehdisadeghidev/croniq/config"
	"github.com/robfig/cron/v3"
)

func Setup() {
	http := resty.New()

	go run(http)
}

func run(http *resty.Client) {
	c := cron.New(cron.WithSeconds())

	for _, job := range config.CronJobs {
		_, _ = c.AddFunc(job.Pattern, func() {
			_, _ = http.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("Accept", "application/json").
				SetAuthToken(config.Token).
				SetBody(job.Data).
				Post(job.Endpoint)
		})
	}

	c.Run()
}

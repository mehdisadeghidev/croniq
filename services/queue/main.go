package queue

import (
	"github.com/go-resty/resty/v2"
	"github.com/mehdisadeghidev/croniq/config"
)

type MainQueue struct {
	Name     string
	Endpoint string
	Data     interface{}
}

var (
	MainChannel    chan MainQueue
	WorkingChannel chan bool
)

func Setup() {
	MainChannel = make(chan MainQueue, 10000)
	WorkingChannel = make(chan bool, 10000)
	http := resty.New()

	for i := 0; i < 10; i++ {
		go run(http)
	}
}

func run(http *resty.Client) {
	for {
		select {
		case c := <-MainChannel:
			WorkingChannel <- true

			_, _ = http.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("Accept", "application/json").
				SetAuthToken(config.Token).
				SetBody(c.Data).
				Post(c.Endpoint)
		}
	}
}

func Shutdown() error {
	if Size() == 0 {
		return nil
	}

	for {
		select {
		case c := <-WorkingChannel:
			if c == false {
				return nil
			}
		}
	}
}

func Size() int {
	return len(MainChannel)
}

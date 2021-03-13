package gin_prometheus_pusher

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type PrometheusConfiguration struct {
	Collectors        *[]prometheus.Collector // User defined Collectors
	Gatherers         *[]prometheus.Gatherer  // User defined Gatherers
	Job               string                  // Job Name
	Address           string                  // Address of the Prometheus Push Gateway
	BasicAuthUser     string                  // Http Basic Auth User
	BasicAuthPassword string                  // Http Basic Auth Password
	AfterRequest      bool                    // Push metrics before request
	BeforeRequest     bool                    // Push metrics after request
}

func Prometheus(config PrometheusConfiguration) gin.HandlerFunc {
	return func(c *gin.Context) {
		wg := &sync.WaitGroup{}
		wg.Add(3)
		pusher := push.New(config.Address, config.Job)
		go configureBasicAuth(config.BasicAuthUser, config.BasicAuthPassword, pusher, wg)
		go attachCollectors(config.Collectors, pusher, wg)
		go attachGatherers(config.Gatherers, pusher, wg)
		wg.Wait()
		if config.BeforeRequest {
			go pushMetrics(pusher)
		}
		c.Next()
		if config.AfterRequest {
			go pushMetrics(pusher)
		}
	}
}

func configureBasicAuth(username string, password string, pusher *push.Pusher, wg *sync.WaitGroup) {
	if username != "" || password != "" {
		pusher.BasicAuth(username, password)
	}
	wg.Done()
}

func attachCollectors(collectors *[]prometheus.Collector, pusher *push.Pusher, wg *sync.WaitGroup) {
	if collectors != nil {
		for _, collector := range *collectors {
			pusher.Collector(collector)
		}
	}
	wg.Done()
}

func attachGatherers(gatherers *[]prometheus.Gatherer, pusher *push.Pusher, wg *sync.WaitGroup) {
	if gatherers != nil {
		for _, gatherer := range *gatherers {
			pusher.Gatherer(gatherer)
		}
	}
	wg.Done()
}

func pushMetrics(pusher *push.Pusher) {
	if err := pusher.Push(); err != nil {
		log.Println(err)
	}
}

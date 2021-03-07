# gin_prometheus_pusher
A Gin Middleware that pushes to a Prometheus Instance instead of scraping it.
Ideal for Serverless Applications

## Index
* [Installation](#installation)
* [Example](#example)

## Installation
```
go get github.com/fourhundredfour/gin_prometheus_pusher
```

## Example
```golang
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    ginprometheuspusher "github.com/fourhundredfour/gin_prometheus_pusher"
)

func main() {
    engine := gin.Default()
	
    var configuration gin_prometheus_pusher.PrometheusConfiguration
    configuration.Address = "http://example.org/metrics"
    configuration.Job = "my_job"
    engine.Use(gin_prometheus_pusher.Prometheus(configuration))

    engine.GET("/", func(ctx *gin.Context) {
        context.JSON(http.StatusOK, nil)
    })
    http.ListenAndServe(":8080", engine)
}
```
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
    "github.com/fourhundredfour/gin_prometheus_pusher"
)

func main() {
    engine := gin.Default()
    engine.Use(gin_prometheus_pusher.Prometheus(&gin_prometheus_pusher.PrometheusConfiguration{
    	Address: "http://example.org/metrics",
	Job: "my_job",
    }))

    engine.GET("/", func(ctx *gin.Context) {
        context.JSON(http.StatusOK, nil)
    })
    http.ListenAndServe(":8080", engine)
}
```

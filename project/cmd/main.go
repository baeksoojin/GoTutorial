package main

import (
	"github.com/go-co-op/gocron"
	"project/pkg/client"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	var p = prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//scheduler
	loc, _ := time.LoadLocation("Asia/Seoul")
	s := gocron.NewScheduler(loc)
	s.Every(1).Day().At("7:30;9:00;11:00;13:00;15:00;17:00;19:00;2:30").Do(client.GetData)
	s.StartAsync()
	e.Logger.Fatal(e.Start(":1323"))
}

var task = func() {}

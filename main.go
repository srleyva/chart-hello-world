package main

import (
	"fmt"
	"github.com/0neSe7en/echo-prometheus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

// Resource in the REST API
type Resource interface {
	Get(c echo.Context) error
	Put(c echo.Context) error
	Post(c echo.Context) error
	Delete(c echo.Context) error
}

// Vending machine
type Vending struct {
	coins     int
	chocolate int
	cost      int
}

// Get the information about the vending machine
func (v *Vending) Get(c echo.Context) error {
	if v.chocolate == 0 {
		return c.String(http.StatusOK, "Out of chocolate please refill")
	}
	return c.String(http.StatusOK, fmt.Sprintf("There are %v pieces of chocolate left.", v.chocolate))
}

// Post will refill the vending machine with the chocolate
func (v *Vending) Post(c echo.Context) error {
	refill, err := strconv.Atoi(c.Param("amount"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please fill with chocolate")
	}
	v.chocolate += refill
	return c.String(http.StatusOK, "Vending machine refilled")
}

// Put will return the chocolate in exchange for the coins
func (v *Vending) Put(c echo.Context) error {
	coins, err := strconv.Atoi(c.Param("amount"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Thats not a coin! No chocolate for you!")
	}
	customerAmount := coins / v.cost
	if customerAmount > v.chocolate {
		return echo.NewHTTPError(http.StatusBadRequest, "There is not enough chocolate to give you")
	}
	v.chocolate -= customerAmount
	v.coins += coins
	return c.String(http.StatusOK, fmt.Sprintf("You now have %v pieces of chocolate", customerAmount))
}

func main() {
	e := echo.New()
	machine := Vending{
		coins:     0,
		chocolate: 10,
		cost:      1,
	}

	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMetricWithConfig(echoprometheus.PrometheusConfig{Namespace: "vending"}))

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.GET("/vending", machine.Get)
	e.POST("/vending/:amount", machine.Post)
	e.PUT("/vending/:amount", machine.Put)

	e.Logger.Fatal(e.Start(":80"))
}

package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"gopkg.in/ugjka/go-tz.v2/tz"
	"net/http"
)

type location struct {
	Lat float64 `json:"lat" form:"name" query:"lat"`
	Lon float64 `json:"lon" form:"email" query:"lon"`
}

func main() {
	e := echo.New()
	e.HideBanner = true
	e.GET("/", index)
	e.POST("/tz/", getTz)
	e.Logger.Fatal(e.Start(":31415"))
}

// return a list of timezone strings
func getTz(c echo.Context) (err error) {
	// get location from json post data
	loc := new(location)
	if err = c.Bind(loc); err != nil {
		return
	}

	// get timezone from coordinates
	zone, err := tz.GetZone(tz.Point{
		Lon: loc.Lon, Lat: loc.Lat,
	})
	if err != nil {
		return
	}

	// prepare json response
	bytes, err := json.Marshal(&zone)
	if err != nil {
		return
	}

	return c.JSONBlob(http.StatusOK, bytes)
}

// return a default string
func index(c echo.Context) (err error) {
	return c.String(http.StatusOK, "")
}

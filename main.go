package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<form action="/" method="post">
				<label for="fname">First name:</label><br>
				<input type="text" id="fname" name="fname" value="John"><br>
				<label for="lname">Last name:</label><br>
				<input type="text" id="lname" name="lname" value="Doe"><br><br>
				<input type="submit" value="Submit">
			</form>
		`)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.POST("/", func(c echo.Context) error {
		fname := c.FormValue("fname")
		lname := c.FormValue("lname")
		return c.HTML(http.StatusOK, "<h1>Hello "+fname+" "+lname+"!</h1>")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

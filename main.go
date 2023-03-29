package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const RedirectsFile = "_redirects"

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		value, err := getKey("/")
		if err != nil || value == "" {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return c.Redirect(302, value)
	})
	e.GET("/:key", func(c echo.Context) error {
		value, err := getKey("/" + c.ParamValues()[0])
		if err != nil || value == "" {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return c.Redirect(302, value)
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func getKey(key string) (string, error) {
	file, err := os.Open(RedirectsFile)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, key+" ") {
			return text[len(key+" "):], scanner.Err()
		}
	}

	return "", scanner.Err()
}

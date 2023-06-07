package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type testModel struct {
	Keyword string `query:"keyword"`
}

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		var input testModel
		input.Keyword = c.QueryParam("keyword")

		return c.String(http.StatusOK, fmt.Sprintf("Input value: %v", input.Keyword))
	})

	e.Logger.Fatal(e.Start(":8001"))
}

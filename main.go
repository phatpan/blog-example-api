package main

import (
	"net/http"

	"github.com/labstack/echo"
	blog "github.com/phatpan/blog-example-api/blog"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/blog", blog.GetBlogList)
	e.Logger.Fatal(e.Start(":1323"))
}

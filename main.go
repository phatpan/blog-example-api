package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/phatpan/blog-example-api/blog"
	"github.com/phatpan/blog-example-api/profile"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.PUT, echo.GET},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/blog", blog.GetBlog)
	e.PUT("/blog/clap/:number", blog.UpdateClap)
	e.GET("/profile", profile.GetProfile)
	e.POST("/profile", profile.UpdateProfile)
	e.Logger.Fatal(e.Start(":1323"))
}

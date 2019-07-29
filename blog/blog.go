package blog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
	Clap    int64  `json:"clap"`
}

func GetBlog(c echo.Context) error {
	file, _ := ioutil.ReadFile("blog/data.json")
	data := blog{}
	_ = json.Unmarshal([]byte(file), &data)

	return c.JSON(http.StatusOK, data)
}

func UpdateClap(c echo.Context) error {
	number, _ := strconv.ParseInt(c.Param("number"), 10, 64)
	file, _ := ioutil.ReadFile("blog/data.json")

	data := blog{}

	_ = json.Unmarshal([]byte(file), &data)

	data.Clap = number
	file, _ = json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("blog/data.json", file, 0644)
	return c.JSON(http.StatusOK, data)
}

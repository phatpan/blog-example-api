package blog

import (
	"net/http"

	"github.com/labstack/echo"
)

type response struct {
	result []blog `json:"result"`
}

type blog struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Content  string `json:"errorMsg"`
	Clapping int    `json:"clapping"`
}

func GetBlogList(c echo.Context) error {
	response := response{
		result: []blog{
			blog{
				ID:       1,
				Name:     "On The Subject Of Subjects (in RxJS)",
				Content:  "Subjects in RxJS are often misunderstood. Because they allow you to imperatively push values into an observable stream, people tend to abuse Subjects when they’re not quite sure how to make an Observable out of something. The pattern looks a little like this",
				Clapping: 1,
			},
			blog{
				ID:       2,
				Name:     "On The Subject Of Subjects (in RxJS)",
				Content:  "Subjects in RxJS are often misunderstood. Because they allow you to imperatively push values into an observable stream, people tend to abuse Subjects when they’re not quite sure how to make an Observable out of something. The pattern looks a little like this",
				Clapping: 1,
			},
		},
	}
	return c.JSON(http.StatusOK, response)
}

package blog

import (
	"net/http"

	"github.com/labstack/echo"
)

type blog struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Content  string `json:"content"`
	Clapping int    `json:"clapping"`
}

func GetBlogList(c echo.Context) error {
	return c.JSON(http.StatusOK, blog{
		ID:       1,
		Title:    "Understanding rxjs BehaviorSubject, ReplaySubject and AsyncSubject",
		Url:      "https://miro.medium.com/max/1000/1*3xWKyEoCO-6DY2DkQKFpvw.jpeg",
		Content:  "Subjects are used for multicasting Observables. This means that Subjects will make sure each subscription gets the exact same value as the Observable execution is shared among the subscribers. You can do this using the Subject class. But rxjs offers different types of Subjects, namely: BehaviorSubject, ReplaySubject and AsyncSubject.",
		Clapping: 0,
	})
}

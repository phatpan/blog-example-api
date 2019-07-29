package profile

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type profile struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

func GetProfile(c echo.Context) error {
	file, _ := ioutil.ReadFile("profile/data.json")
	data := profile{}
	_ = json.Unmarshal([]byte(file), &data)

	return c.JSON(http.StatusOK, data)
}

func UpdateProfile(c echo.Context) error {
	profile := new(profile)
	if err := c.Bind(profile); err != nil {
		return err
	}

	if profile.Name == "azumo" { // Simulate case return error
		return c.JSON(http.StatusBadRequest, "Could not update "+profile.Name)
	}

	file, _ := json.MarshalIndent(profile, "", " ")

	_ = ioutil.WriteFile("profile/data.json", file, 0644)
	return c.JSON(http.StatusOK, profile)
}

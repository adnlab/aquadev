package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	//return c.String(http.StatusOK, id)

	return c.JSON(http.StatusOK, Item{
		Success: true,
		Data: User{
			Name: c.Param(name: "Name"),
		},
	})
}

func main() {
	e := echo.New()
	e.GET("/users/:id", getUser)

	// Start echo server after all code have been read
	e.Logger.Fatal(e.Start(":1323"))
}

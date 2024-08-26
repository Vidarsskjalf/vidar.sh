package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "layout", map[string]interface{}{
		"title":   "Home",
		"Content": "home",
	})
}

func ProjectsHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "layout", map[string]interface{}{
		"title":   "Projects",
		"Content": "projects",
	})
}

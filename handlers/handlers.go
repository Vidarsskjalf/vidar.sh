package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "layout", map[string]interface{}{
		"Title":   "Home",
		"Content": "home",
	})
}

func ProjectsHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "layout", map[string]interface{}{
		"Title":   "Projects",
		"Content": "projects",
	})
}

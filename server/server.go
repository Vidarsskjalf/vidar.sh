package server

import (
    "html/template"
    "io"
    "net/http"
    "log"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type Server struct {
    *echo.Echo
}

type TemplateRenderer struct {
    templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func StartServer() {
    e := echo.New()
    // Render templates
    renderer := &TemplateRenderer{
        templates: template.Must(template.ParseGlob("template/*.html")),
    }
    e.Renderer = renderer
    // Server static files
    e.Static("/", "static")
    // Middleware
    e.Use(middleware.Logger())

  if err := e.Start(":8080"); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}

func (s *Server) Start() error {
    s.Logger.Info("Server is starting on port 8080...")
    return s.Echo.Start(":8080")
}

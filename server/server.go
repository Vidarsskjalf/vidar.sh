package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"vidar.sh/handlers"
)

type Server struct {
	*echo.Echo
	port string
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewServer(port string) *Server {
	e := echo.New()

	// Render templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Server static files
	e.Static("/static", "static")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.HomeHandler)
	e.GET("/projects", handlers.ProjectsHandler)

	return &Server{Echo: e, port: port}
}

func (s *Server) Start() error {
	s.Logger.Infof("Server is starting on port %s...", s.port)
	return s.Echo.Start(s.port)
}

package main

import (
	"errors"
	"fmt"
	"io"
	"test/weekday/routes"
	"text/template"

	"github.com/labstack/echo"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func MakeTemplate(e *echo.Echo) {
	templates := make(map[string]*template.Template)
	templates["weekday.html"] = template.Must(template.ParseFiles("views/weekday.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}
}

func main() {
	fmt.Println("Start program calculate weekday from date")
	// fs := http.FileServer(http.Dir("assets/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	e := echo.New()

	MakeTemplate(e)
	rt := routes.Routers{e}
	rt.GetRouter()

	e.Start(":80")
}
package main

import (
	"html/template"
	"io"

	"github.com/aadi-1024/notes/pkg/models"
	"github.com/labstack/echo/v4"
)

type Template struct {
	temp *template.Template
}

// implement the Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.temp.ExecuteTemplate(w, name, data)
}

func SetupRouter(e *echo.Echo) {
	files := []string{"./templates/home.page.gohtml", "./templates/base.layout.gohtml"}
	t := &Template{
		temp: template.Must(template.ParseFiles(files...)),
	}
	e.Renderer = t

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Hi\n")
	})
	e.GET("/", func(c echo.Context) error {
		if app.debug {
			t = &Template{
				temp: template.Must(template.ParseFiles(files...)),
			}
			e.Renderer = t
		}
		return c.Render(200, "home.page.gohtml", map[string]any{
			"LoggedIn": true,
			"Notes": []models.Note{
				{
					Id:    1,
					Title: "Title 1",
					Text:  "Sample",
				},
				{
					Id:    2,
					Title: "Title 2",
					Text:  "Sample 2",
				},
			},
		})
	})
}

package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Contact struct {
	Name  string
	Email string
}

var contacts = map[string][]*Contact{
	"Contacts": {
		{Name: "Misael", Email: "misa@example.com"},
		{Name: "Fernanda", Email: "fernanda@example.com"},
		{Name: "Marco", Email: "m@example.com"},
	},
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index", contacts)
}

func addContact(c echo.Context) error {
	name := c.Request().PostFormValue("name")
	email := c.Request().PostFormValue("email")
	return c.Render(http.StatusOK, "contact-list-info", &Contact{Name: name, Email: email})
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
	}
	e.Renderer = t
	e.GET("/", Hello)
	e.POST("/add-contact", addContact)
	e.Logger.Fatal(e.Start(":80"))
}

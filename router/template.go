package router

import (
	"html/template"
	"io"
	"github.com/labstack/echo"
	"errors"
)

var templates = make(map[string]*template.Template)

func NewTemplates() *TemplateRegistry {
	templates["index.html"] = template.Must(template.ParseFiles("public/views/index.html", "public/views/layout.html"))
	templates["edit.html"] = template.Must(template.ParseFiles("public/views/edit.html", "public/views/layout.html"))

	return &TemplateRegistry { templates: templates }
}

type TemplateRegistry struct {
    templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
	  err := errors.New("Template not found -> " + name)
	  return err
	}
	return tmpl.ExecuteTemplate(w, "layout.html", data)
}

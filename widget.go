package main

import (
	"html/template"
	"io"
)

type Widget struct {
	Name     string                 `json:"name"`
	Position string                 `json:"position"`
	Content  map[string]interface{} `json:"content"`
}

func (w *Widget) Render(f io.Writer, s *Site) error {
	basename := w.Name + ".html.tpl"
	path := "themes/" + s.Theme + "/widgets/"

	t, err := template.New(basename).
		Funcs(template.FuncMap{"safeHTML": safeHTML}).
		ParseFiles(path + basename)

	if err != nil {
		return err
	}

	return t.Execute(f, w)
}

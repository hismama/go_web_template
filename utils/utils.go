package utils

import (
	"go_web_template/templates"
	"log"
	"net/http"
)

func Check(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	tpl, err := templates.Tpl.Clone()
	if err != nil {
		return err
	}
	tpl, err = tpl.ParseFiles(
		"templates/pages/" + tmpl,
	)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tpl.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		return err
	}

	return nil
}

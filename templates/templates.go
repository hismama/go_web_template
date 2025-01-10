package templates

import (
	"html/template"
)

var Tpl *template.Template

func Init() {
	Tpl = template.Must(template.ParseGlob("templates/helpers/*.gohtml"))
	Tpl = template.Must(Tpl.ParseGlob("templates/base/*.gohtml"))
}

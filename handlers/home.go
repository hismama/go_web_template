package handlers

import (
	"go_web_template/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := utils.RenderTemplate(w, "index.gohtml", nil)
	utils.Check(err, w)
}

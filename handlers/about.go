package handlers

import (
	"go_web_template/utils"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderTemplate(w, "about.gohtml", nil)
	utils.Check(err, w)
}

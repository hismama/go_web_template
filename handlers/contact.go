package handlers

import (
	"go_web_template/utils"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderTemplate(w, "contact.gohtml", nil)
	utils.Check(err, w)
}

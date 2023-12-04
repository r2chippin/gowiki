package controller

import (
	"fmt"
	model "gowiki/models"
	"net/http"
)

// ViewHandler sad
func ViewHandler(w http.ResponseWriter, r *http.Request, td string) {
	title := r.URL.Path[len("/view/"):]
	p, _ := model.LoadPage(title, td)
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<div>%s</div>", p.Title, p.Body)
	if err != nil {
		panic(err)
	}
}

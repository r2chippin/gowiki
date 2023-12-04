package controller

import (
	"fmt"
	model "gowiki/models"
	"net/http"
)

// ViewHandler sad
func ViewHandler(w http.ResponseWriter, r *http.Request, td string) {
	title := r.URL.Path[len("/view/"):]
	p, err := model.LoadPage(title, td)
	if err != nil {
		p = &model.Page{Title: title}
	}
	_, err = fmt.Fprintf(w, "<h1>%s</h1>\n<div>%s</div>", p.Title, p.Body)
	if err != nil {
		panic(err)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request, td string) {
	title := r.URL.Path[len("/view/"):]
	p, err := model.LoadPage(title, td)
	if err != nil {
		p = &model.Page{Title: title}
	}
	_, err = fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
	if err != nil {
		panic(err)
	}
}

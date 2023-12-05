package controller

import (
	model "gowiki/models"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmplName string, p *model.Page, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, tmplName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, td string, tmpl *template.Template) {
	title := r.URL.Path[len("/view/"):]
	p, err := model.LoadPage(title, td)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p, tmpl)
}

func EditHandler(w http.ResponseWriter, r *http.Request, td string, tmpl *template.Template) {
	title := r.URL.Path[len("/view/"):]
	p, err := model.LoadPage(title, td)
	if err != nil {
		p = &model.Page{Title: title}
	}
	renderTemplate(w, "edit", p, tmpl)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, td string) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &model.Page{Title: title, Body: []byte(body)}
	err := p.Save(td)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

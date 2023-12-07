package controller

import (
	"errors"
	model "gowiki/models"
	"html/template"
	"net/http"
	"regexp"
)

func renderTemplate(w http.ResponseWriter, tmplName string, p *model.Page, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, tmplName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTitle(w http.ResponseWriter, r *http.Request, vp *regexp.Regexp) (string, error) {
	m := vp.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string), vp *regexp.Regexp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := vp.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, td string, tmpl *template.Template, vp *regexp.Regexp) {
	title, err := getTitle(w, r, vp)
	if err != nil {
		return
	}
	p, err := model.LoadPage(title, td)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p, tmpl)
}

func EditHandler(w http.ResponseWriter, r *http.Request, td string, tmpl *template.Template, vp *regexp.Regexp) {
	title, err := getTitle(w, r, vp)
	if err != nil {
		return
	}
	p, err := model.LoadPage(title, td)
	if err != nil {
		p = &model.Page{Title: title}
	}
	renderTemplate(w, "edit", p, tmpl)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, td string, vp *regexp.Regexp) {
	title, err := getTitle(w, r, vp)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &model.Page{Title: title, Body: []byte(body)}
	err = p.Save(td)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gowiki/models"
	"net/http"
)

func ViewPageHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")
		p, err := models.LoadPage(title, db)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/wiki/new/"+title)
			return
		}
		ctx.HTML(http.StatusOK, "view.tmpl", gin.H{
			"Title":   p.Title,
			"Content": p.Content,
		})
	}
}

func NewWikiHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")
		p, err := models.LoadPage(title, db)
		if err != nil {
			p = &models.Page{Title: title}
			ctx.HTML(http.StatusOK, "new.tmpl", gin.H{
				"Title": p.Title,
			})
		} else {
			ctx.Redirect(http.StatusFound, "/wiki/view/"+title)
		}
	}
}

func EditPageHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")
		p, err := models.LoadPage(title, db)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/wiki/view/"+title)
			return
		}
		ctx.HTML(http.StatusOK, "edit.tmpl", gin.H{
			"Title":   p.Title,
			"Content": p.Content,
		})
	}
}

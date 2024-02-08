package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gowiki/models"
	"net/http"
)

func CreateWikiHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")
		content := ctx.PostForm("content")
		var p1 models.Page
		db.Unscoped().Where("title = ?", title).Find(&p1)
		db.Unscoped().Delete(&p1)

		p := &models.Page{
			Title:   title,
			Content: []byte(content),
		}
		err := p.Save(db)
		if err != nil {
			ctx.HTML(http.StatusOK, "fail.tmpl", gin.H{
				"Title": title,
			})
			return
		}
		ctx.HTML(http.StatusOK, "success.tmpl", gin.H{
			"Title": title,
		})
	}
}

func UpdateWikiHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")
		content := ctx.PostForm("content")
		p, err := models.LoadPage(title, db)
		if err != nil {
			return
		}
		err = p.Update([]byte(content), db)
		if err != nil {
			ctx.HTML(http.StatusOK, "fail.tmpl", gin.H{
				"Title": title,
			})
			return
		}
		ctx.HTML(http.StatusOK, "success.tmpl", gin.H{
			"Title": title,
		})
	}
}

func DeleteWikiHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")
		p, err := models.LoadPage(title, db)
		if err != nil {
			return
		}
		err = p.Delete(db)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "fail.tmpl", gin.H{
				"Title": title,
			})
			return
		}
		ctx.HTML(http.StatusOK, "success.tmpl", gin.H{
			"Title": title,
		})
	}
}

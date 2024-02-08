package routes

import (
	"github.com/gin-gonic/gin"
	"gowiki/api"
	"gowiki/config"
)

func NewRouter(cfg config.Config) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("tmpl/*")

	db, _ := config.InitDB(cfg)

	wiki := r.Group("/wiki")
	{
		wiki.GET("/ping", api.PingPong())

		wiki.GET("/view/:title", api.ViewPageHandler(db))
		wiki.GET("/new/:title", api.NewWikiHandler(db))
		wiki.GET("/edit/:title", api.EditPageHandler(db))

		wiki.POST("/create/:title", api.CreateWikiHandler(db))
		wiki.PATCH("/update/:title", api.UpdateWikiHandler(db))
		wiki.DELETE("/delete/:title", api.DeleteWikiHandler(db))
	}

	return r
}

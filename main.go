package main

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.Static("/css", "./css")

	r.HTMLRender = createMyRenderer()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about", gin.H{
			"Title": "关于",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func createMyRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles(
		"index",
		"templates/layouts/default.html",
		"templates/includes/head.html",
		"templates/includes/header.html",
		"templates/includes/footer.html",
		"templates/pages/index.html",
	)
	r.AddFromFiles(
		"about",
		"templates/layouts/page.html",
		"templates/includes/head.html",
		"templates/includes/header.html",
		"templates/includes/footer.html",
		"templates/pages/about.html",
	)
	return r
}

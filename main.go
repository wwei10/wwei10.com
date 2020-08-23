package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"

	"github.com/wwei10/wwei10.com/parser"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.Static("/css", "./css")
	r.Static("/favicon.ico", "./assets/favicon.ico")

	r.HTMLRender = createMyRenderer()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default", gin.H{
			"Content": "",
		})
	})

	// Generate pages from directory pages.
	pages := parser.GetPagesMapFromDir("./templates/pages")
	fmt.Println(pages)

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page", gin.H{
			"Title":   "关于",
			"Content": template.HTML(blackfriday.Run([]byte(pages["/about/"].Content))),
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
		"default",
		"templates/layouts/default.html",
		"templates/includes/head.html",
		"templates/includes/header.html",
		"templates/includes/footer.html",
		"templates/pages/index.html",
	)
	r.AddFromFiles(
		"page",
		"templates/layouts/page.html",
		"templates/includes/head.html",
		"templates/includes/header.html",
		"templates/includes/footer.html",
	)
	return r
}

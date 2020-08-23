package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"

	"github.com/wwei10/wwei10.com/parser"
)

// Renders timeline of the blog.
func timeline(c *gin.Context) {
	posts := parser.GetPagesFromDir("./posts")
	c.HTML(http.StatusOK, "index", gin.H{
		"Posts": posts,
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.Static("/css", "./css")
	r.Static("/favicon.ico", "./assets/favicon.ico")

	r.HTMLRender = createMyRenderer()

	// Generate feed.
	r.GET("/", timeline)

	// Generate Posts
	r.GET("/posts/:postname", func(c *gin.Context) {
		postname := c.Param("postname")
		posts := parser.GetPagesFromDir("./posts")
		for _, post := range posts {
			if strings.Contains(post.Permalink, postname) {
				c.HTML(http.StatusOK, "page", gin.H{
					"Title":   post.Title,
					"Content": template.HTML(blackfriday.Run([]byte(post.Content))),
				})
			}
		}
	})

	// Generate pages from directory pages.
	pages := parser.GetPagesMapFromDir("./templates/pages")
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page", gin.H{
			"Title":   pages["/about/"].Title,
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
		"index",
		"templates/layouts/index.html",
		"templates/includes/head.html",
		"templates/includes/header.html",
		"templates/includes/footer.html",
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

package main

import (
	"net/http"
	"strings"

	"github.com/Depado/bfchroma"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"

	"github.com/wwei10/wwei10.com/parser"
)

func timelineAPI(c *gin.Context) {
	var posts = parser.GetPagesFromDir("./posts")
	for i := range posts {
		posts[i].Content = string(blackfriday.Run(
			[]byte(posts[i].Content),
			blackfriday.WithRenderer(
				// See options here:
				// https://github.com/alecthomas/chroma/tree/master/styles
				bfchroma.NewRenderer(bfchroma.Style("dracula")),
			),
		))
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func searchAPI(c *gin.Context) {
	var posts = parser.GetPagesFromDir("./posts")
	link := c.Param("link")
	for i := range posts {
		if strings.Contains(posts[i].Permalink, link) {
			posts[i].Content = string(blackfriday.Run(
				[]byte(posts[i].Content),
				blackfriday.WithRenderer(
					// See options here:
					// https://github.com/alecthomas/chroma/tree/master/styles
					bfchroma.NewRenderer(bfchroma.Style("dracula")),
				),
			))
			c.JSON(http.StatusOK, gin.H{
				"post": posts[i],
			})
			return
		}
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	// APIs.
	v1 := r.Group("/api/v1")
	{
		v1.GET("/timeline", timelineAPI)
		v1.GET("/search/:link", searchAPI)
	}

	// Serve react.
	r.Use(static.Serve("/", static.LocalFile("./app/build", false)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./app/build/index.html")
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

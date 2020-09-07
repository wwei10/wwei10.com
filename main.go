package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/Depado/bfchroma"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wwei10/wwei10.com/ginzap"
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
	var path = nil
	if gin.IsDebugging() {
		path = "wwei10.com.log"
	} else {
		path = "/root/wwei10.com.log"
	}
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"wwei10.com.log"},
	}.Build()

	// Add a ginzap middleware which:
	//  - Logs all requests, like a combined access and error log.
	//  - Logs to stdout.
	//  - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

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

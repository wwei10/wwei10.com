package main

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/Depado/bfchroma"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/blackfriday/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wwei10/wwei10.com/counter"
	"github.com/wwei10/wwei10.com/ginzap"
	"github.com/wwei10/wwei10.com/parser"
)

var db, err = sql.Open("sqlite3", "./stats.db")

func timelineAPI(c *gin.Context) {
	var posts = parser.GetPagesFromDir("./posts")

	// Increment page view counter.
	counter.UpdateDB(*db, "timeline")

	// Aggregate all posts.
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
		"posts":     posts,
		"totalView": counter.GetTotalViews(*db),
	})
}

func searchAPI(c *gin.Context) {
	var posts = parser.GetPagesFromDir("./posts")
	link := c.Param("link")

	// Increment page view counter.
	counter.UpdateDB(*db, link)

	// Search for relevant post.
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
				"post":      posts[i],
				"postView":  counter.GetStats(*db, link),
				"totalView": counter.GetTotalViews(*db),
			})
			return
		}
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	var path = "/root/wwei10.com.log"
	if gin.IsDebugging() {
		path = "wwei10.com.log"
	}
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{path},
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

	// Serve react in release mode.
	// In debugging mode, use npm start.
	if !gin.IsDebugging() {
		r.Use(static.Serve("/", static.LocalFile("./app/build", false)))
		r.NoRoute(func(c *gin.Context) {
			c.File("./app/build/index.html")
		})
	}
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

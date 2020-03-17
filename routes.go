// routes.go

package main

import (
	"github.com/msrshahrukh100/Golang-webdev-reference/handlers"
)

func initializeRoutes() {

  // Handle the index route
  router.GET("/", handlers.ShowIndexPage)
  router.GET("/article/view/:article_id", handlers.GetArticle)
}
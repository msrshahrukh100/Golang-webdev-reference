// handlers.article.go

package handlers

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "github.com/msrshahrukh100/Golang-webdev-reference/models"
)


func render(c *gin.Context, data gin.H, templateName string) {

  switch c.Request.Header.Get("Accept") {
  case "application/json":
    // Respond with JSON
    c.JSON(http.StatusOK, data["payload"])
  case "application/xml":
    // Respond with XML
    c.XML(http.StatusOK, data["payload"])
  default:
    // Respond with HTML
    c.HTML(http.StatusOK, templateName, data)
  }

}


func ShowIndexPage(c *gin.Context) {
  articles := models.GetAllArticles()

  // Call the HTML method of the Context to render a template
  render(c, gin.H{
    "title":   "Home Page",
    "payload": articles}, "index.html")

}

func GetArticle(c *gin.Context) {
  // Check if the article ID is valid
  if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
    // Check if the article exists
    if article, err := models.GetArticleByID(articleID); err == nil {
      // Call the HTML method of the Context to render a template

      render(c, gin.H{
        "title":   article.Title,
        "payload": article}, "article.html")

    } else {
      // If the article is not found, abort with an error
      c.AbortWithError(http.StatusNotFound, err)
    }

  } else {
    // If an invalid article ID is specified in the URL, abort with an error
    c.AbortWithStatus(http.StatusNotFound)
  }
}
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserHostShow(c *gin.Context) {
	c.HTML(http.StatusOK, "host.html", gin.H{
		"title": "userHost",
	})
}

func AdminHostShow(c *gin.Context) {
	c.HTML(http.StatusOK, "construct.html", gin.H{
		"title": "AdminHost",
	})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s AdminUser) Index(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
package controllers

import (
	"github.com/gin-gonic/gin"
)

func (s AdminUser) Index(c *gin.Context) {


	s.display(c)
}

func (s AdminUser) Show(c *gin.Context) {


	s.display(c)
}
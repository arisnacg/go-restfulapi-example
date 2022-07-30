package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "hello from controller!"})
}

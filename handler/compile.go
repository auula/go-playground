package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/higker/go-playground/response"
)

// Compile provide function for compile go code
func Compile(ctx *gin.Context) {
	form := ctx.PostForm("body")
	if len(form) < 14 {
		response.FailJson(ctx, "code context error.")
		return
	}

}

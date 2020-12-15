package router

import (
	"github.com/gin-gonic/gin"
	"github.com/higker/go-playground/handler"
)

// mappingCompiler mapping router for go compiler
func mappingCompiler(r *gin.Engine) {
	//TODO
	r.POST("/compile", handler.Compile)
}

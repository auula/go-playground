// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/15 - 7:58 下午 - UTC/GMT+08:00

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	linkView = func(c *gin.Context) {
		c.HTML(http.StatusOK, "link.html", gin.H{})
	}
	editorView = func(c *gin.Context) {
		c.HTML(http.StatusOK, "editor.html", gin.H{})
	}
)

// mappingView mapping router for view
func mappingView(r *gin.Engine) {
	r.StaticFS("/static", http.Dir("static"))
	//加载模板
	r.LoadHTMLGlob("template/*")
	r.GET("/", editorView)
	r.GET("/index.html", editorView)
	r.GET("/link", linkView)
}

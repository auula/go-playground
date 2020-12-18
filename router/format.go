// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/17 - 9:19 下午 - UTC/GMT+08:00

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/higker/go-playground/handler"
)

// mappingFormat mapping router for code format api
func mappingFormat(r *gin.Engine) {
	r.GET("/format", handler.FormatCode)
}

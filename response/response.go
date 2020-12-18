// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/18 - 5:25 下午 - UTC/GMT+08:00

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	OkJson = func(ctx *gin.Context, data interface{}) {
		ctx.JSON(http.StatusOK, NewResponse(200, data))
	}
	FailJson = func(ctx *gin.Context, data interface{}) {
		ctx.JSON(http.StatusOK, NewResponse(500, data))
	}
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, data interface{}) *Response {
	return &Response{Code: code, Data: data}
}

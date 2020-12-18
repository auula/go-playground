// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/17 - 9:19 下午 - UTC/GMT+08:00

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/higker/go-playground/response"
)

var formatCode gin.HandlerFunc

func FormatCode(ctx *gin.Context) {
	response.OkJson(ctx, response.NewResponse(200, "TEST"))
}

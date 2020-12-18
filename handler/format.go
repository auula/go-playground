// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/17 - 9:19 下午 - UTC/GMT+08:00

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/higker/go-playground/response"
	"go/format"
)

type fmtResponse struct {
	Body  string
	Error string
}

func FormatCode(ctx *gin.Context) {
	form := ctx.PostForm("body")
	source, err := format.Source([]byte(form))
	if err != nil {
		response.FailJson(ctx, fmtResponse{Error: err.Error()})
		return
	}
	response.OkJson(ctx, fmtResponse{Body: string(source)})
}

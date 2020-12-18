// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/17 - 9:19 下午 - UTC/GMT+08:00

package handler

import (
	"fmt"
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
	fmt.Println(form)
	if len(form) <= 14 {
		response.FailJson(ctx, fmtResponse{Error: "code body error."})
		return
	}
	source, err := format.Source([]byte(form))
	if err != nil {
		response.FailJson(ctx, fmtResponse{Error: err.Error()})
		return
	}
	response.OkJson(ctx, fmtResponse{Body: string(source)})
}

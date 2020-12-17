// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/15 - 7:49 下午 - UTC/GMT+08:00

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/higker/go-playground/config"
)

var (
	root *gin.Engine
)

func init() {
	gin.SetMode(gin.DebugMode)
	root = gin.New()
	root.Use(gin.Logger(), gin.Recovery())
}

//Start 运行Web
func Start() {
	mappingView(root)
	mappingCompiler(root)
	mappingFormat(root)
	_ = root.Run(config.Port)
}

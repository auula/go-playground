// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/15 - 7:49 下午 - UTC/GMT+08:00

package router

import "github.com/gin-gonic/gin"

const (
	port = ":8080"
)

//Start 运行Web
func Start() {
	r := gin.Default()
	mappingView(r)
	mappingCompiler(r)
	_ = r.Run(port)
}

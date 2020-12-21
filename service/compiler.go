// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/21 - 9:43 下午 - UTC/GMT+08:00

package service

import (
	"os"
	"os/exec"
	"sync"
)

var (
	// docker里面test.go文件指针
	source *os.File
	lock   sync.Mutex
)

func Builder(sourceCode []byte) (string, error) {
	var dockerImageId string
	file, err := os.OpenFile("workspace/test.go", os.O_WRONLY, 0666)
	if err != nil {
		return "", err
	}
	lock.Lock()
	defer lock.Unlock()
	source = file
	_, err = source.Write(sourceCode)
	source.Close()
	if err != nil {
		return "", err
	}
	source.Close()
	cmd := ""
	exec.Command("docker", "build")
	return "", nil
}

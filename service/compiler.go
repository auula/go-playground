// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/12/21 - 9:43 下午 - UTC/GMT+08:00

package service

import (
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

var (
	// docker里面test.go文件指针
	source *os.File
	lock   sync.Mutex
)

//func Builder(sourceCode []byte) (string, error) {
//	var dockerImageId string
//	file, err := os.OpenFile("workspace/test.go", os.O_WRONLY, 0666)
//	if err != nil {
//		return "", err
//	}
//	lock.Lock()
//	defer lock.Unlock()
//	source = file
//	_, err = source.Write(sourceCode)
//	source.Close()
//	if err != nil {
//		return "", err
//	}
//	source.Close()
//	cmd := ""
//	exec.Command("docker", "build")
//	return "", nil
//}

// 这个函数最好是低权限用户运行，不然会出现远程执行漏洞，高级漏洞
func Builder(sourceCode []byte) (string, error) {
	var dockerImageId string
	file, err := os.OpenFile("workspace/test.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
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
	dockerImageId = strconv.Itoa(int(time.Now().UnixNano()))
	cmd := "cd workspace && docker build -t goplay:" + dockerImageId
	output, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	cmd = "docker run goplay:" + dockerImageId
	output, err = exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	//cmd = "docker rmi -f goplay:" + dockerImageId
	//output, err = exec.Command("/bin/bash", "-c", cmd).Output()
	//if err != nil {
	//	return "", err
	//}
	return string(output), nil
}

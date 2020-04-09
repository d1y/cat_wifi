// 测试某个指令是否存在
// link: https://gist.github.com/miguelmota/ed4ec562b8cd1781e7b20151b37de8a0

package main

import "os/exec"

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	isError := err == nil
	return isError
}

var adbIsInstall bool

func init() {
	// 测试 `adb` 指令是否安装
	adbIsInstall = commandExists("adb")
}
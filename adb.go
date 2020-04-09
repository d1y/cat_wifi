package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 运行`adb`命令
func runADB(cmds string, ids ...string) (string, error) {
	var cmdArr []string
	tempArr := strings.Split(cmds, " ")
	if len(ids) >= 1 {
		cmdArr = append(cmdArr, "-s", ids[0])
	}
	cmdArr = append(cmdArr,tempArr...)
	out := exec.Command("adb", cmdArr...)
	resp, err := out.Output()
	result :=  strings.TrimSpace(string(resp))
	return result, err
}

// 打开手机设置
func openWifiSetting() bool {
	_, err := runADB("shell am start -a android.intent.action.MAIN -n com.android.settings/.wifi.WifiSettings")
	if err != nil {
		return false
	}
	return true
}

// 截图拉取图片
func pullImage() string {
	// 未root的手机上, 只有 `/data/local/tmp` 目录权限读写
	phoneTempPath := filepath.Join(tempPath, "./screen.png")
	localPath := filepath.Join(os.TempDir(),"./screen.png")
	cmd := fmt.Sprintf("shell %s -p %s", screencapBin, phoneTempPath)
	pullCmd := fmt.Sprintf("pull %s %s", phoneTempPath, localPath)
	rmScreenFileCmd := fmt.Sprintf("shell %s %s", rmBin, phoneTempPath)
	runADB(cmd)
	runADB(pullCmd)
	runADB(rmScreenFileCmd)
	return localPath
}

// 点亮屏幕
func turnOnScreen() bool {
	_, err := runADB("shell input keyevent 26")
	if err != nil {
		return false
	}
	return true
}

// 判断屏幕是否点亮
func screenIsLight () bool {
	resp, _ := runADB("shell dumpsys display | grep \"mScreenState\"")
	searchKey := "mScreenState="
	index := strings.Index(resp, searchKey)
	if (index >= 0) {
		word := resp[len(searchKey):]
		if strings.Index(word, "OFF") >= 0 {
			return false
		} else if strings.Index(word, "ON") >= 0 {
			return true
		}
	}
	return false
}

// 设备是否锁定 https://stackoverflow.com/a/56553363/10272586
func checkUnlocked() bool {
	resp, _ := runADB("shell dumpsys window | grep \"mDreamingLockscreen\"")
	searchKey := "mDreamingLockscreen="
	index := strings.Index(resp, searchKey)
	if (index >= 0) {
		len := index + len(searchKey)
		sp := strings.TrimSpace(resp[len:len+5])
		if (sp == "false") {
			return false
		} else if (sp == "true") {
			return true
		}
	}
	return true
}

// 测试是否 `root`
func testRoot() bool {
	outRoot := exec.Command("adb", "root")
	resp, err := outRoot.Output()
	if err != nil {
		return false
	}
	str := strings.TrimSpace(string(resp))
	if str == asRootErrorMsg {
		fmt.Println("比对msg失败")
	}
	return true
}

// 获取设备
func getDevices() []devicesBody {
	var body []devicesBody
	outDevices := exec.Command("adb", "devices", "-l")
	resp, err := outDevices.Output()
	if err != nil {
		// unknownThrowMsg
		return body
	}
	str := strings.TrimSpace(string(resp))
	split := strings.Split(str, "\n")

	if len(split) <= 1 {
		// noDevicesMsg
		return body
	}
	splitx := split[1:]
	arr := uncodeDevices(splitx)
	body = arr
	// getDevicesSucessMsg
	return body
}

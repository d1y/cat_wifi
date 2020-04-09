// WIFI:T:WPA;S:FAST_TEST;P:6666;;

package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	log.Println("正在查找手机")
	lists := getDevices()
	if len(lists) >= 1 {
		current := lists[0]
		log.Println("目前找到", len(lists), "台设备")
		log.Println("===================")
		log.Println("该设备型号: ", current.model)
		log.Println("自动连接该设备...")
		log.Println("开启屏幕中...")
		turnOnScreen()
		isUnlocked := checkUnlocked()
		if isUnlocked {
			log.Println("该手机已锁定,请先解锁手机")
			ticker := time.NewTicker(1 * time.Second)
			i := 0
			for range ticker.C {
				testUnlock := checkUnlocked()
				if !testUnlock {
					log.Println("解锁手机成功, 用时", i, "秒")
					ticker.Stop()
					break
				}
				i++
			}
		} else {
			log.Println("手机已解锁")
		}
		log.Println("自动打开Wifi设置")
		isLight := screenIsLight()
		if !isLight {
			turnOnScreen()
		}
		openWifiSetting()
		log.Println("将需要解密的wifi打开分享弹窗")
		fmt.Print("我已打开分享弹窗[Y/N]: ")
		var isShareModal string
		fmt.Scan(&isShareModal)
		toUpperShareModal := strings.ToUpper(isShareModal)
		if toUpperShareModal == "Y" {
			log.Println("输入正确")
			tempPath := pullImage()
			log.Println("正在截图, 临时文件: ", tempPath)
			body, err := qrcodeToWifiBody(tempPath)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Wifi名称: ", body.SID)
				log.Println("Wifi密码: ", body.PWD)
			}
		}
		log.Println("===================")
	}
}

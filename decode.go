// 解密和格式化字符串code

package main

import (
	"fmt"
	"strings"
)

type devicesBody struct {
	id          string // 设备id
	usb         string
	product     string
	model       string // 设备型号
	device      string // 设备
	transportID string
}

// 格式化手机信息
func uncodeDevices(codes []string) []devicesBody {
	var arr []devicesBody
	for _, value := range codes {
		var body devicesBody
		items := strings.Split(value, " ")
		for _, item := range items {
			newItem := strings.TrimSpace(item)
			if len(newItem) >= 1 && !(newItem == "device") {
				if strings.Contains(newItem, ":") {
					Kv := strings.Split(newItem, ":")
					key := Kv[0]
					value := Kv[1]
					switch key {
					case "usb":
						body.usb = value
						fallthrough
					case "product":
						body.product = value
						fallthrough
					case "model":
						body.model = value
						fallthrough
					case "device":
						body.device = value
					case "transport_id":
						body.transportID = value
					}
				} else {
					body.id = newItem
				}
			}
		}
		arr = append(arr, body)
	}
	return arr
}

type wifiBody struct {
	isFormat bool   // 格式是否对
	Type     string // 加密类型
	SID      string // 名称
	PWD      string // 密码
}

// 格式化 wifi
func decodeWifi(code string) wifiBody {
	var Body wifiBody
	var x1 = "WIFI:"
	var x2 = ";;"
	var isFormat = strings.Index(code, x1) == 0 && len(code)-strings.LastIndex(code, x2) == 2
	Body.isFormat = isFormat
	if isFormat {
		var firstTempIndex = len(x1)
		var lastTempIndex = (len(code) - 2)
		var newCode = code[firstTempIndex:lastTempIndex]
		// T:WPA;S:FAST_TEST;P:6666
		var newCodeArr = strings.Split(newCode, ";")
		Body.Type = newCodeArr[0][2:]
		Body.SID = newCodeArr[1][2:]
		Body.PWD = newCodeArr[2][2:]
	} else {
		fmt.Println("未知错误")
	}
	return Body
}

// 格式化 `.conf` 文件
func decodeRootWifiPassword(code string) string {
	var str string
	if len(code) == 0 {
		str = catWifiPassword()
	} else {
		str = code
	}
	arr := strings.Split(str, "\n")
	var networkArr []string
	var keyword string = "network={"
	var index int
	index = -1
	for idx, item := range arr {
		item = strings.TrimSpace(item)
		if item == keyword {
			index = idx
		} else {
			if item == "}" {
				// TODO
				index = -1
			} else {
				index = -1
			}
		}
	}
	return str
}

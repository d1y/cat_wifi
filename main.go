// WIFI:T:WPA;S:FAST_TEST;P:6666;;

package main

import (
	"fmt"
	"os"
	"strings"
)

type wifiBody struct {
	isFormat bool   // 格式是否对
	Type     string // 加密类型
	SID      string // 名称
	PWD      string // 密码
}

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

func main() {
	var args = os.Args

	if len(args) >= 2 {
		var body = decodeWifi(args[1])
		if (body.isFormat) {
			fmt.Println("wifi加密类型: ", body.Type)
			fmt.Println("wifi名称: ", body.SID)
			fmt.Println("wifi密码: ", body.PWD)
		} else {
			fmt.Println("参数错误")
		}
	} else {
		fmt.Println("参数错误")
	}

}

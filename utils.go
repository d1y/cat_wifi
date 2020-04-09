package main

import (
	"errors"
	"os"
	"github.com/tuotoo/qrcode"
)

// https://studygolang.com/articles/24617
// 读取截图的二维码
func qrcodeToWifiBody(path string) (wifiBody, error) {
	var body wifiBody
	tempFile := "/Users/kozo4/Desktop/screen.png"
	fi, err := os.Open(tempFile)
	if err != nil {
		return body, errors.New("读取文件错误")
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		return body, errors.New("读取二维码失败")
	}
	ctx := qrmatrix.Content
	face := decodeWifi(ctx)
	if len(ctx) == 0 {
		return face, errors.New("读取二维码失败")
	}
	return face, nil
}

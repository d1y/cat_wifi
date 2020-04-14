package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	e "errors"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func qrcodeToWifiBody(path string) (wifiBody, error) {
	var body wifiBody
	// open and decode image file
	file, err := os.Open(path)
	if err != nil {
		return body, e.New("文件读取失败")
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return body, e.New("图片解析失败")
	}
	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return body, e.New("图片解析失败")
	}
	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return body, e.New("二维码编码失败")
	}
	// TODO
	spf := fmt.Sprintf("%s", result)
	body = decodeWifi(spf)
	return body, nil
}

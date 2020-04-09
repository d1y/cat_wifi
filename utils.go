package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func qrcodeToWifiBody(path string) wifiBody {
	// open and decode image file
	file, _ := os.Open(path)
	img, _, _ := image.Decode(file)
	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)
	spf := fmt.Sprintf("%s", result)
	body := decodeWifi(spf)
	return body
}

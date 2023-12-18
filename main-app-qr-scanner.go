package main

import (
	"fmt"
	"image"

	/*
		Depending of variations of format you might read
		you need to include those image types below else
		you would get error. I have taken into account JPEG
		and PNG
	*/
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"go.mercari.io/go-emv-code/mpm"
)

func main() {
	// open and decode image file
	file, _ := os.Open(os.Args[1])
	img, _, _ := image.Decode(file)

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	fmt.Println(result.GetText())

	dst, err := mpm.Decode([]byte(result.GetText()))
	fmt.Println(dst.PayloadFormatIndicator, err)
	fmt.Println(dst.MerchantAccountInformation)
	fmt.Println(dst.MerchantAccountInformation[0].Tag)
	fmt.Println(dst.MerchantAccountInformation[0].Length)
	fmt.Println(dst.MerchantAccountInformation[0].Value)
	// [{15 31 **999166**999166****M0004203432} {26 37 0018NG.COM.NIBSSPLC.QR0111S0003850946}]
	// in this the first object carries merchant number starting with M
	// and the second object carries the sub merchant number starting with S
}

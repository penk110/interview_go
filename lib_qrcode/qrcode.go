package main

import (
	"image/color"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	// var png []byte
	// var err error
	// png, err = qrcode.Encode("https://github.com/zyphub", qrcode.Medium, 256)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("png: %v\n", string(png))

	err := qrcode.WriteColorFile("https://github.com/zyphub", qrcode.Medium, 256, color.Black, color.White, "qr.png")
	if err != nil {
		panic(err)
	}

}

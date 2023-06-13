package qrcode

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func QrCodeGen(t string, filename string) error {
	qrCode, _ := qr.Encode(t, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 2000, 2000)
	file, err := os.Create("./static/qr/" + filename)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

	return err
}

func GenQR() {
	// This is a text to encode
	t := "Hello"
	// file to read from
	filename := "qrcode2.png"
	QrCodeGen(t, filename)
}

package qrcode

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"

	// github. com liyue201 GOQR
	"github.com/liyue201/goqr"
)

func ScanQRCode(path string) {

	// Reads imgdata from a file.
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Decodes an image.
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return
	}

	// Recognize QR Codes.
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return
	}

	// Prints a list of QR Codes.
	for _, qrCode := range qrCodes {
		fmt.Printf("qrCode text: %s\n", qrCode.Payload)
	}
	return
}

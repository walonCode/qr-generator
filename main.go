package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main(){
	args := os.Args
	if len(args) < 3 || len(args) > 3 {
		log.Fatalf("usage: qr-generator <link> <filename.png>")
	}

	link := args[1]
	filename := args[2]

	err := qrcode.WriteColorFile(link, qrcode.High, 256, color.White, color.Black, filename)
	if err != nil {
		log.Fatalf("Error creating qrcode: %v\n",err)
	}

	fmt.Printf("qrcode generated and saved to %s\n", filename)
}
package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: missing required arguments.")
		fmt.Fprintln(os.Stderr, "Usage: qr-generator <link> <filename.png>")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Arguments:")
		fmt.Fprintln(os.Stderr, "  <link>          The URL or text to encode into the QR code")
		fmt.Fprintln(os.Stderr, "  <filename.png>  Output file path (must end in .png)")
		os.Exit(1)
	}

	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, "Error: missing argument <filename.png>.")
		fmt.Fprintln(os.Stderr, "Usage: qr-generator <link> <filename.png>")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, "Error: too many arguments — expected exactly 2.")
		fmt.Fprintln(os.Stderr, "Usage: qr-generator <link> <filename.png>")
		os.Exit(1)
	}

	link := args[0]
	filename := args[1]

	err := qrcode.WriteColorFile(link, qrcode.High, 256, color.White, color.Black, filename)
	if err != nil {
		log.Fatalf("Error: could not generate QR code: %v\n", err)
	}

	fmt.Printf("QR code generated and saved to %s\n", filename)
}
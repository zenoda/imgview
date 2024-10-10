package main

import (
	"github.com/zenoda/imgview"
	"image/jpeg"
	"os"
)

func main() {
	f, err := os.Open("pic.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	imgview.Show(img)
}

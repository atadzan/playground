package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	imgPath := "./image/sddefault.jpg"
	if err := cropImg(imgPath); err != nil {
		log.Fatal(err.Error())
		return
	}
}

func cropImg(imgPath string) error {
	file, err := os.Open(imgPath)
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	file.Close()

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	div := width / 16
	extraSpace := height - 9*div
	cropHeight := extraSpace / 2
	if cropHeight <= 0 {
		return nil
	}

	x0, y0 := 0, cropHeight
	x1, y1 := img.Bounds().Max.X, img.Bounds().Max.Y-cropHeight

	croppedImg := img.(*image.YCbCr).SubImage(image.Rect(x0, y0, x1, y1))

	croppedFile, err := os.Create(imgPath)
	if err != nil {
		return err
	}

	defer croppedFile.Close()

	if err = jpeg.Encode(croppedFile, croppedImg, nil); err != nil {
		return err
	}
	return nil
}

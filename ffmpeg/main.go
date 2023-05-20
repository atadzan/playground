package main

import (
	"github.com/atadzan/playground/fastHttp"
	image "github.com/atadzan/playground/ffmpeg/convert"
	"github.com/atadzan/playground/minio"
	"io"
	"log"
	"os"
)

func main() {
	storage, err := minio.InitMinio()
	if err != nil {
		log.Println("error while initializing minio storage. Error: ", err.Error())
		return
	}
	url := "https://vid.puffyan.us/vi/V7jVbEcnz8o/maxres.jpg"
	response, err := fastHttp.SendRequest(url, fastHttp.CreateGetRequest(url))
	if err != nil {
		log.Println("error while getting img from net.Error: ", err.Error())
	}
	convertedImgPath := "../assets/test-image.webp"

	if err = image.ConvertJpgToWebpFromResponseBody(response.Body(), convertedImgPath); err != nil {
		log.Println("can't convert jpg to webp. Error: ", err.Error())
		return
	}
	file, err := os.Open(convertedImgPath)
	if err != nil {
		log.Println("can't open file. Error: ", err.Error())
		return
	}
	imgBody, err := io.ReadAll(file)
	if err != nil {
		log.Println("can't read img body. Error: ", err.Error())
		return
	}
	bucket := "image"
	imgStoragePath := "test/test-img.webp"
	if err = storage.UploadConvertedImage(bucket, imgStoragePath, imgBody); err != nil {
		log.Println("error while uploading converted img. Error: ", err.Error())
		return
	}
}

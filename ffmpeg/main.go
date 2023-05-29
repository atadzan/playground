package main

import (
	"fmt"
	image "github.com/atadzan/playground/ffmpeg/convert"
	"log"
	"time"
)

//func main() {
//	storage, err := minio.InitMinio()
//	if err != nil {
//		log.Println("error while initializing minio storage. Error: ", err.Error())
//		return
//	}
//	url := "https://vid.puffyan.us/vi/V7jVbEcnz8o/maxres.jpg"
//	response, err := fastHttp.SendRequest(url, fastHttp.CreateGetRequest(url))
//	if err != nil {
//		log.Println("error while getting img from net.Error: ", err.Error())
//	}
//	convertedImgPath := "../assets/test-image.webp"
//
//	if err = image.ConvertJpgToWebpFromResponseBody(response.Body(), convertedImgPath); err != nil {
//		log.Println("can't convert jpg to webp. Error: ", err.Error())
//		return
//	}
//	file, err := os.Open(convertedImgPath)
//	if err != nil {
//		log.Println("can't open file. Error: ", err.Error())
//		return
//	}
//	imgBody, err := io.ReadAll(file)
//	if err != nil {
//		log.Println("can't read img body. Error: ", err.Error())
//		return
//	}
//	bucket := "image"
//	imgStoragePath := "test/test-img.webp"
//	if err = storage.UploadConvertedImage(bucket, imgStoragePath, imgBody); err != nil {
//		log.Println("error while uploading converted img. Error: ", err.Error())
//		return
//	}
//}

func main() {
	start := time.Now()
	if err := image.EncodeToHEVCGood(); err != nil {
		log.Println("error occured. Error: ", err.Error())
		return
	}
	duration := time.Since(start)

	fmt.Println("Executed time(good): ", duration)

	start1 := time.Now()
	if err := image.EncodeToHEVCDefault(); err != nil {
		log.Println("error occured. Error: ", err.Error())
		return
	}
	duration1 := time.Since(start1)

	fmt.Println("Executed time(default): ", duration1)
}

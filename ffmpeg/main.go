package main

import (
	"bytes"
	"fmt"
	"github.com/atadzan/playground/fastHttp"
	image "github.com/atadzan/playground/ffmpeg/convert"
	"github.com/atadzan/playground/minio"
	"io"
	"log"
	"os"
)

func DownloadFile(path string, body io.Reader) {
	file, err := os.Create(path)
	if err != nil {
		log.Println("error while creating file. Error: ", err.Error())
	}
	_, err = io.Copy(file, body)
	if err != nil {
		log.Println("error while copying. Error: ", err.Error())
	}

}

func main() {
	storage, err := minio.InitMinio()
	if err != nil {
		log.Println("error while initializing minio storage. Error: ", err.Error())
		return
	}
	localPath := "../assets/test-5.jpg"
	url := "https://vid.puffyan.us/vi/V7jVbEcnz8o/maxres.jpg"
	respons, err := fastHttp.SendRequest(url, fastHttp.CreateGetRequest(url))

	DownloadFile(localPath, bytes.NewReader(respons.Body()))
	//err = storage.UploadImage(path, response)
	if err != nil {
		log.Println("error while uploading img", err.Error())
	}
	presignedUrl, err := storage.GenerateUploadUrl("videos", "test/test-5.webp")
	if err != nil {
		return
	}
	reader, err := image.ConvertJpgToWebpWithOutput(localPath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	////if err = storage.UploadConvertedImage(path, reader); err != nil {
	////	log.Println("error while uploading converted img. Error: ", err.Error())
	////	return
	////}
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	response, err := fastHttp.SendRequest(presignedUrl, fastHttp.CreatePutRequest(presignedUrl, buf.Bytes()))
	if response.StatusCode() != 200 {
		log.Println("error while sending request.Error: ", err.Error())
	}

	//fasthttp.Request{}
	fmt.Println("Success")
	//image.ConvertJpgToWebp()
}

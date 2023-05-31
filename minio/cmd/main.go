package main

import (
	"fmt"
	"github.com/atadzan/playground/minio"
	"log"
)

func main() {
	mc, err := minio.InitMinio()
	if err != nil {
		log.Println("Error while init", err.Error())
		return
	}
	if err = mc.CreateBucket(); err != nil {
		log.Println("Error while creating", err.Error())
		return
	}
	fmt.Println("Finish")
}

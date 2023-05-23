package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
)

func main() {
	conn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{
		User: "admin",
		Pass: "pass",
	})
	if err != nil {
		log.Println("can't connect to tarantool. Error: ", err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Connected")

	//trnl := customCrud.NewTarantool(conn)
	//
	//if err != nil {
	//	return
	//}
	//
	//code := resp.Code
	//data := resp.Data
	//fmt.Println("Data:", data, "\n Code:", code)
}

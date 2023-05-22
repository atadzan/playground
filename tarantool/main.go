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

	// Insert data
	//resp, err := conn.Insert("tester", []interface{}{4, "Roxy", 2023})
	//if err != nil {
	//	log.Println("can't insert input data. Error:", err.Error())
	//	return
	//}

	//// Get data
	//resp, err := conn.Select("tester", "primary", 0, 1, tarantool.IterEq, []interface{}{4})
	//if err != nil {
	//	log.Println("can't select data. Error: ", err.Error())
	//	return
	//}

	// Update data
	//resp, err := conn.Update("tester", "primary", []interface{}{4}, []interface{}{[]interface{}{"+", 2, 3}})
	//if err != nil {
	//	log.Println("can't update data. Error: ", err.Error())
	//	return
	//}

	// Replace data
	//resp, err := conn.Replace("tester", []interface{}{4, "new band", 2011})
	//if err != nil {
	//	log.Println("can't replace data. Error: ", err.Error())
	//	return
	//}

	// Upsert data
	//resp, err := conn.Upsert("tester", []interface{}{4, "Another band", 2000}, []interface{}{[]interface{}{"+", 2, 5}})
	//if err != nil {
	//	log.Println("can't upsert. Error:", err.Error())
	//	return
	//}

	// Delete data
	resp, err := conn.Delete("tester", "primary", []interface{}{3})
	if err != nil {
		log.Println("can't delete. Error: ", err.Error())
		return
	}
	code := resp.Code
	data := resp.Data
	fmt.Println("Data:", data, "\n Code:", code)
}

package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	aaa [2]bool // смещение байтов: 0
	ccc [2]bool // смещение байтов: 8
	bbb int32   // смещение байтов: 4
}

type Foo2 struct {

	//aaa int32 // 4
	ccc int64 // 4
	bbb int32 // 4
	ddd bool  //

}

func main() {
	//x := Foo{}
	x := Foo2{}

	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Alignof(x))
}

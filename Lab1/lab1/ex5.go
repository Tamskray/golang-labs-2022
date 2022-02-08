package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Синонимы целых типов\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")

	//Задание.
	//1. Определить разрядность ОС

	var x uint32
	a := unsafe.Pointer(&x)
	if unsafe.Sizeof(a) == 8 {
		fmt.Println("x64")
	} else if unsafe.Sizeof(a) == 4 {
		fmt.Println("x32")
	}

	/*
		const PtrSize = 32 << uintptr(^uintptr(0)>>63)
		fmt.Println(runtime.GOOS, runtime.GOARCH)
		fmt.Println(strconv.IntSize, PtrSize)
	*/
}

package main

import (
	"fmt"
)

var tmpX int

const tmpY int = 20

func main3() {
	fmt.Println("Test3")

	type Box struct {
		width  int
		height int
	}

	box := Box{30, 20}
	fmt.Println("Box W=", box.width, " H=", box.height, " >>", box)

	fmt.Println("before X=", tmpX, " , Y=", tmpY)
	tmpX = 40
	//	tmpY = 20	// const error! java final
	fmt.Println("after X=", tmpX, " , Y=", tmpY)

	//	const (
	//		one   = 0
	//		two   = 1
	//		three = 2
	//		four  = 3
	//	)

	//	const (
	//		one = iota
	//		two
	//		three
	//		four
	//	)

	//	const (
	//		one = iota + 2
	//		two
	//		three
	//		four
	//	)

	type Number int
	const (
		//		_          = iota
		one Number = 5 + iota
		two
		three
		four
	)

	fmt.Println(one, " / ", two, " / ", three, " / ", four)
}

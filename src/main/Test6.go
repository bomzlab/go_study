package main

import (
	"fmt"
	m "parsetest"
)

func main6() {

	v := m.ToInt("3")
	fmt.Println("toInt :: ", v)

	s := m.ToString(33)
	fmt.Println("toString :: ", s)
}

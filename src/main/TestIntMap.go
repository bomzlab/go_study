package main

import (
	"fmt"
	mp "model/maps"
)

func maina() {
	fmt.Println("intMap test")

	m := mp.NewIntMap()
	m.Add(3, "good")

	key := 3
	fmt.Println("key ", key, " , value = ", m.Get(key))

}

package main

import (
	"fmt"
	bmap "model/maps"
)

func main() {
	//	m := bmap.NewSyncMap()
	//	fmt.Println("count %d", m.Count())
	//
	//	m.Set("king", "kong")
	//	fmt.Println("count %d", m.Count())

	m := bmap.NewBomzStyleMapInstance()
	printMap(m)

	m.Add("good", "bye")
	printMap(m)

	fmt.Println("get :: ", m.Get("good"))
	fmt.Println("get :: ", m.Get("good1"))
}

func printMap(m bmap.BomzMap) {
	fmt.Println(m, " , count : ", m.Count())
}

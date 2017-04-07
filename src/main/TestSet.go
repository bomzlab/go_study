package main

import (
	"fmt"
	set "model/sets"
)

func main() {
	fmt.Printf("% start\n", "bomzSet")

	var mySet set.BomzSet

	mySet = getSet(true)

	fmt.Println(">>", mySet)

	mySet.Add(3)
	mySet.Add(8)

	for i := 0; i < 10; i++ {
		fmt.Println(i, " = ", mySet.Contains(i))
	}

	fmt.Println("===================", mySet.Count())

	mySet.Remove(8)

	for i := 0; i < 10; i++ {
		fmt.Println(i, " = ", mySet.Contains(i))
	}

	fmt.Println("===================", mySet.Count())

	v := mySet.Values()

	for i, r := range v {
		fmt.Println(">>>>", i, "=", r)
	}
}

func getSet(sync bool) set.BomzSet {
	if sync {
		return set.NewSyncSet()
	} else {
		return set.NewNonSyncSet()
	}
}

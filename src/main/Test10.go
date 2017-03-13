package main

import (
	"fmt"
	bomzComms "model/comms"
	bomzMap "model/maps"
)

func main10() {
	m := bomzMap.New()

	fmt.Println("map : ", m)

	m["king"] = "kong"
	m["hot"] = "dog"

	fmt.Println("map : ", m)

	process(m)
}

func process(g bomzComms.GetMap) {
	fmt.Println("process >> ", g.Get("hot"))
}

package main

import (
	"fmt"
	user "model/user"
)

func main9() {
	info, err := user.New("bomz", "min", 15, "south", "korea")
	fmt.Println("make user :: ", info, " <> ", err)

	if err != nil {
		return
	}

	info.SetAge(30)

	fmt.Println("update age user :: ", info)
}

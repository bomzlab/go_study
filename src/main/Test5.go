package main

import (
	"fmt"
	"strings"
)

func main5() {

	v := func(x, y int) int {
		return x + y
	}

	v(3, 4)

	//	v := func(x, y int) int {
	//		return x + y
	//	}(3, 4)
	fmt.Println("1>", v)

	a := sumFileName(".zip")
	fmt.Println("2>", a("good"))

	fmt.Println("3>", sum1(3, sum2))

	hangulCheck("Hi, 안녕, Good")
}

func sumFileName(prefix string) func(string) string {
	return func(fileName string) string {
		if strings.HasSuffix(fileName, prefix) {
			return fileName
		}
		return fileName + prefix
	}
}

func sum1(y int, f func(int, int) int) int {
	return f(y, 3)
}

func sum2(x, y int) int {
	return x + y
}

func hangulCheck(s string) {
	i := strings.Index(s, "안")
	fmt.Println("index = ", i, " , 본문 = ", s)
}

package main

import (
	"fmt"
	"strconv"
)

func main4() {
	a, b := 3, 5

	fmt.Println("sum : ", sum(a, b))
	fmt.Println("mult : ", mult(a, b))

	x, y := test1(a, b)
	fmt.Println("test1 : ", x, " / ", y)

	toInt1("3a")
	fmt.Println("toInt2 : ", toInt2("4"))

	fmt.Println(point1(&a), " , main = ", a)

	defer1(a, b)
}

func sum(x, y int) int {
	return x + y
}

func mult(x, y int) int {
	return x * y
}

func test1(x, y int) (int, int) {
	return sum(x, y), mult(x, y)
}

func toInt1(v string) {
	if i, err := strconv.Atoi(v); err == nil {
		fmt.Println("value = ", i)
	} else {
		fmt.Println("value ", v, " is error. ", err)
	}
}

func toInt2(v string) (re int) {
	re, _ = strconv.Atoi(v)
	fmt.Println("toInt2 method : ", re)
	return
}

func point1(x *int) string {
	*x += 3
	return "Point1 method = " + strconv.Itoa(*x)
}

func defer1(x, y int) {
	fmt.Println("defer1-1 : ", x, " , ", y)
	v := x + y

	defer fmt.Println("defer1-2 : ", x, " , ", y, " = ", v)

	v = 40
	fmt.Println("defer1-3 : ", x, " , ", y, " = ", v)
}

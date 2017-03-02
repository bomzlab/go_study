package main

import (
	"fmt"
)

const MAX = 30

func main2() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("sum=", sum)

	sum = 8

	switch {
	case sum >= 0 && sum <= 10:
		fmt.Println("0~10 : ", sum)
		fallthrough // 다음 case 실행
	case sum > 10 && sum <= 20:
		fmt.Println("11~20 : ", sum)
	case sum > 20 && sum <= 40:
		fmt.Println("21~40 : ", sum)
	case sum > 40 && sum <= 50:
		fmt.Println("41~50 : ", sum)
	}

	array := []int{3, 4, 5, 6}
	for i := 0; i < len(array); i++ {
		fmt.Println("array ", i, " = ", array[i])
	}
}

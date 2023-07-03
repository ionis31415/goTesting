package main

import (
	"fmt"
)

func calculateDaysBetweenDates(date1, date2 string) int {

	return 0
}

func update(px *int) {
	*px = 100
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor * factor
	}
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	x := 10

	fmt.Println("start")

	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
	update(&x)
	fmt.Println(x)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}

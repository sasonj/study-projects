// package main

// import (
// 	"fmt"
// )

// func main2() {
// 	var number int
// 	fmt.Print("enter a number you'd like to check")
// 	fmt.Scan(&number)
// 	if number%2 == 1 {
// 		fmt.Print(number, "нечетное")
// 	} else {
// 		fmt.Print(number, "четное")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func familyName(fname string) {
// 	fmt.Println("Hello", fname)
// }

// func main() {
// 	a := "Artem"
// 	familyName(a)
// 	familyName("Jenny")
// 	familyName("Anja")
// }

package main

import (
	"fmt"
)

func myFunction(x int, y int) int {
	sum := x + y
	return sum
}

func main() {
	var x, y int
	fmt.Println("введите числа")
	fmt.Scan(&x, &y)

	sum := myFunction(x, y)
	fmt.Println("Сумма чисел равна", sum)
}

// package main

// import (
// 	"fmt"
// )

//  func main() {
// 	const Pi float64 = 3.14
// 	fmt.Println("Чтобы узнать длину и площадь окружности введите радиус")

// 	var ra float64
// 	var length, square float64

// 	fmt.Scan(&ra)
// 	length = 2 * Pi * ra
// 	fmt.Printf("Длина окружности - %.2f\n", length)

// 	fmt.Scan(&square)
// 	square = Pi * ra * ra
// 	fmt.Printf("Площадь окружности - %.2f\n", square)

// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	var input string
// 	buf := bufio.NewReader(os.Stdin)

// 	input, err := buf.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}

// 	var cntSym, cntWrd, cntStr uint
// 	for i, sym := range input {
// 		fmt.Print(i, string(sym))
// 		cntSym++
// 		if sym == ' ' {
// 			cntWrd++
// 		}
// 		if sym == '\n' {
// 			cntStr++
// 		}

// 	}
// 	fmt.Println(cntSym-2, cntWrd+1, cntStr)
// }

package main

import (
	"fmt"
)

func mainv() {

}

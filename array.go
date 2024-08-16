Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example,
the program belongs to the main package.

Line 2: import ("fmt") lets us import files included in the fmt package.

Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text.
In our example it will output "Hello World!".

int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
string - stores text, such as "Hello World". String values are surrounded by double quotes
bool- stores values with two states: true or false

Syntax
var variablename type = value
		package main

	import ("fmt")

	func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane" //type is inferred
	x := 2 //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	}

	Variable Declaration Without Initial Value
In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:

Example
		package main
		import ("fmt")

		func main() {
		var a string
		var b int
		var c bool

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		}

		Go Multiple Variable Declaration
In Go, it is possible to declare multiple variables in the same line.

Example
This example shows how to declare multiple variables in the same line:

			package main
			import ("fmt")

			func main() {
			var a, b, c, d int = 1, 3, 5, 7

			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
			}

			Go Variable Declaration in a Block
Multiple variable declarations can also be grouped together into a block for greater readability:

Example
	package main
	import ("fmt")

	func main() {
	var (
		a int
		b int = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	}

	A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Go variable naming rules:

A variable name must start with a letter or an underscore character (_)
A variable name cannot start with a digit
A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
Variable names are case-sensitive (age, Age and AGE are three different variables)
There is no limit on the length of the variable name
A variable name cannot contain spaces
The variable name cannot be any Go keywords

Example
If we want to print the arguments in new lines, we need to use \n.

		package main
		import ("fmt")

		func main() {
		var i,j string = "Hello","World"

		fmt.Print(i, "\n")
		fmt.Print(j, "\n")
		}
		Result:

		Hello
		World
		Tip: \n creates new lines.
				func main() {
					var i,j string = "Hello","World"

					fmt.Print(i, "\n",j)
				}

If we want to add a space between string arguments, we need to use " ":

			package main
			import ("fmt")

			func main() {
			var i,j string = "Hello","World"

			fmt.Print(i, " ", j)
			}
			Result:

			Hello World
			Print() inserts a space between the arguments if neither are strings:

			// The Printf() Function&
			String Data Type
The string data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:

Go Arrays
Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

Declare an Array
In Go, there are two ways to declare an array:

1. With the var keyword:
Syntax
	var array_name = [length]datatype{values} // here length is defined

or

	var array_name = [...]datatype{values} // here length is inferred
2. With the := sign:
Syntax
	array_name := [length]datatype{values} // here length is defined

or

	array_name := [...]datatype{values} // here length is inferred

	Note: The length specifies the number of elements to store in the array.
	In Go, arrays have a fixed length. The length of the array is either defined by
	a number or is inferred (means that the compiler decides the length of the array, based on the number of values).

	Access Elements of an Array
You can access a specific array element by referring to the index number.

In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

Example
This example shows how to access the first and third elements in the prices array:

		package main
		import ("fmt")

		func main() {
		prices := [3]int{10,20,30}

		fmt.Println(prices[0])
		fmt.Println(prices[2])
		}
Result:

10
30

Change Elements of an Array
You can also change the value of a specific array element by referring to the index number.

Example
This example shows how to change the value of the third element in the prices array:

		package main
		import ("fmt")

		func main() {
		prices := [3]int{10,20,30}

		prices[2] = 50
		fmt.Println(prices)
		}
		Result:

[10 20 50]

Initialize Only Specific Elements
It is possible to initialize only specific elements in an array.

Example
This example initializes only the second and third elements of the array:

package main
import ("fmt")

func main() {
  arr1 := [5]int{1:10,2:40}

  fmt.Println(arr1)
}
Result:

[0 10 40 0 0]
Example Explained
The array above has 5 elements.

1:10 means: assign 10 to array index 1 (second element).
2:40 means: assign 40 to array index 2 (third element).

Create a Slice With []datatype{values}
Syntax
slice_name := []datatype{values}
A common way of declaring a slice is like this:

myslice := []int{}
The code above declares an empty slice of 0 length and 0 capacity.

To initialize the slice during declaration, use this:

myslice := []int{1,2,3}
The code above declares a slice of integers of length 3 and also the capacity of 3.

In Go, there are two functions that can be used to return the length and capacity of a slice:

len() function - returns the length of the slice (the number of elements in the slice)
cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
Example
This example shows how to create slices using the []datatype{values} format:

		package main
		import ("fmt")

		func main() {
		myslice1 := []int{}
		fmt.Println(len(myslice1))
		fmt.Println(cap(myslice1))
		fmt.Println(myslice1)

		myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
		fmt.Println(len(myslice2))
		fmt.Println(cap(myslice2))
		fmt.Println(myslice2)
		}
		Result:

		0
		0
		[]
		4
		4
		[Go Slices Are Powerful]

		Create a Slice From an Array
You can create a slice by slicing an array:

Syntax
var myarray = [length]datatype{values} // An array
myslice := myarray[start:end] // A slice made from the array
Example
This example shows how to create a slice from an array:

		package main
		import ("fmt")

		func main() {
		arr1 := [6]int{10, 11, 12, 13, 14,15}
		myslice := arr1[2:4]

		fmt.Printf("myslice = %v\n", myslice)
		fmt.Printf("length = %d\n", len(myslice))
		fmt.Printf("capacity = %d\n", cap(myslice))
		}
		Result:

		myslice = [12 13]
		length = 2
		capacity = 4
In the example above myslice is a slice with length 2. It is made from arr1 which is an array with length 6.

The slice starts from the third element of the array which has value 12
(remember that array indexes start at 0. That means that [0] is the first element,
	 [1] is the second element, etc.). The slice can grow to the end of the array.
	  This means that the capacity of the slice is 4.

If myslice started from element 0, the slice capacity would be 6.
Create a Slice With The make() Function
The make() function can also be used to create a slice.

Syntax
slice_name := make([]type, length, capacity)
Note: If the capacity parameter is not defined, it will be equal to length.

Example
This example shows how to create slices using the make() function:

package main
import ("fmt")

func main() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
}
Result:

myslice1 = [0 0 0 0 0]
length = 5
capacity = 10
myslice2 = [0 0 0 0 0]
length = 5
capacity = 5

Append Elements To a Slice
You can append elements to the end of a slice using the append()function:

Syntax
slice_name = append(slice_name, element1, element2, ...)
Example
This example shows how to append elements to the end of a slice:

package main
import ("fmt")

func main() {
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))
}
Result:

myslice1 = [1 2 3 4 5 6]
length = 6
capacity = 6
myslice1 = [1 2 3 4 5 6 20 21]
length = 8
// capacity = 12

Append One Slice To Another Slice
To append all the elements of one slice to another slice, use the append()function:

Syntax
slice3 = append(slice1, slice2...)
Note: The '...' after slice2 is necessary when appending the elements of one slice to another.

Example
This example shows how to append one slice to another slice:

		package main
		import ("fmt")

		func main() {
		myslice1 := []int{1,2,3}
		myslice2 := []int{4,5,6}
		myslice3 := append(myslice1, myslice2...)

		fmt.Printf("myslice3=%v\n", myslice3)
		fmt.Printf("length=%d\n", len(myslice3))
		fmt.Printf("capacity=%d\n", cap(myslice3))
		}
Result:

myslice3=[1 2 3 4 5 6]
length=6
capacity=6

Change The Length of a Slice
Unlike arrays, it is possible to change the length of a slice.

Example
This example shows how to change the length of a slice:

package main

import (
	"fmt"
)

func main() {
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	arr2 := make([]int, 4)
	copy(arr2, arr1[1:5]) // Slice array
	myslice1 := arr2
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}

// Result:

// myslice1 = [10 11 12 13]
// length = 4
// capacity = 5
// // myslice1 = [10 11]
// length = 2
// capacity = 5
// myslice1 = [10 11 20 21 22 23]
// length = 6
// capacity = 10


ogical Operators
Logical operators are used to determine the logic between variables or values:

Operator	Name	Description	Example	Try it
&& 	Logical and	Returns true if both statements are true	x < 5 &&  x < 10	
|| 	Logical or	Returns true if one of the statements is true	x < 5 || x < 4	
!	Logical not	Reverse the result, returns false if the result is true	!(x < 5 && x < 10)	

Operator	Name	Example	Try it
==	Equal to	x == y	
!=	Not equal	x != y	
>	Greater than	x > y	
<	Less than	x < y	
>=	Greater than or equal to	x >= y	
<=	Less than or equal to	x <= y

Go has the following conditional statements:

Use if to specify a block of code to be executed, if a specified condition is true
Use else to specify a block of code to be executed, if the same condition is false
Use else if to specify a new condition to test, if the first condition is false
Use switch to specify many alternative blocks of code to be executed

he Nested if Statement
You can have if statements inside if statements, this is called a nested if.

Syntax
if condition1 {
   // code to be executed if condition1 is true
  if condition2 {
     // code to be executed if both condition1 and condition2 are true
  }
}
Example
This example shows how to use nested if statements:

	package main
	import ("fmt")

	func main8() {
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
		fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}
}

Single-Case switch Example
The example below uses a weekday number to calculate the weekday name:

Example
		package main
		import ("fmt")

		func main() {
		day := 4

		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		}
		}
Result:

Thursday


The default Keyword
The default keyword specifies some code to run if there is no case match:

Example
		package main
		import ("fmt")

		func main() {
		day := 8

		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Not a weekday")
		}
		}
Result:

Not a weekday

	for Loop Examples
	Example 1
	This example will print the numbers from 0 to 4:

	package main
	import ("fmt")

	func main() {
	for i:=0; i < 5; i++ {
	fmt.Println(i)
	}
	}
Result:

0
1
2
3
4

	Example 2
	This example counts to 100 by tens: 

	package main
	import ("fmt")

	func main() {
	for i:=0; i <= 100; i+=10 {
		fmt.Println(i)
	}
	}
Result:

0
10
20
30
40
50
60
70
80
90
100

The continue Statement
The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

Example
This example skips the value of 3:

		package main
		import ("fmt")

		func main() {
		for i:=0; i < 5; i++ {
			if i == 3 {
			continue
			}
		fmt.Println(i)
		}
		}
Result:

0
1
2
4

The break Statement
The break statement is used to break/terminate the loop execution.

Example
This example breaks out of the loop when i is equal to 3:

	package main
	import ("fmt")

	func main() {
	for i:=0; i < 5; i++ {
		if i == 3 {
		break
		}
	fmt.Println(i)
	}
	}
Result:

0
1
2


Nested Loops
It is possible to place a loop inside another loop.

Here, the "inner loop" will be executed one time for each iteration of the "outer loop":

Example
	package main
	import ("fmt")

	func main() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i:=0; i < len(adj); i++ {
		for j:=0; j < len(fruits); j++ {
		fmt.Println(adj[i],fruits[j])
		}
	}
	}
Result:

big apple
big orange
big banana
tasty apple
tasty orange
tasty banana


Example
This example uses range to iterate over an array and print both the indexes and the values at each 
(idx stores the index, val stores the value):

	package main
	import ("fmt")

	func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
	fmt.Printf("%v\t%v\n", idx, val)
	}
	}
Result:

0      apple
1      orange
2      banana

Here, we want to omit the indexes (idx stores the index, val stores the value):

	package main
	import ("fmt")

	func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
	}
Result:

apple
orange
banana

Call a Function
Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

In the example below, we create a function named "myMessage()". The opening curly brace ( { ) indicates the beginning 
	of the function code, and the closing curly brace ( } ) indicates the end of the function. The function 
	outputs "I just got executed!". To call the function, just write its name followed by two parentheses ():

Example
		package main
		import ("fmt")

		func myMessage() {
		fmt.Println("I just got executed!")
		}

		func main() {
		myMessage() // call the function
		}
Result:

I just got executed!
A function can be called multiple times.

Example
		package main
		import ("fmt")

		func myMessage() {
		fmt.Println("I just got executed!")
		}

		func main() {
		myMessage()
		myMessage()
		myMessage()
		}
Result:

I just got executed!
I just got executed!
I just got executed!
Function With Parameter Example
The following example has a function with one parameter (fname) of type string. When the familyName() function is called,
 we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names,
  but an equal last name:

Example
package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}
Result:

Hello Liam Refsnes
Hello Jenny Refsnes
Hello Anja Refsnes
Note: When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter,
 while Liam, Jenny and Anja are arguments.

Multiple Parameters
Inside the function, you can add as many parameters as you want:

Example
package main
import ("fmt")

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
  familyName("Liam", 3)
  familyName("Jenny", 14)
  familyName("Anja", 30)
}
Result:

Hello 3 year old Liam Refsnes
Hello 14 year old Jenny Refsnes
Hello 30 year old Anja Refsnes
Note: When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters,
 and the arguments must be passed in the same order.

 Multiple Return Values
Go functions can also return multiple values.

Example
Here, myFunction() returns one integer (result) and one string (txt1):

package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
}
Result:

10 Hello World!
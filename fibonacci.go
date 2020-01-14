/*
Exercise: Fibonacci closure
Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/

package main

import "fmt"

func fibonacci() func() int {
  n := 0
  n_1 := 0
  n_2 := 0
  return func() int {
    n += 1
    switch n {
			case 0:
				n_1 = 1
				return n_1
				case 1:
				n_2 = 1
			return n_1
			default:
				sum := n_1 + n_2
			n_2 = n_1
			n_1 = sum
			return sum
		}
  }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

package main

import "fmt"

//custom number type
type number int

func main() {
	printEvenOdd(newNumbers(11))
}

//generates new slice of numbers for given limit
func newNumbers(limit int) []number {
	numbers := []number{}
	for i := 0; i < limit; i++ {
		numbers = append(numbers, number(i))
	}
	return numbers
}

//reciever function to check if number is even
func (n number) isEven() bool {
	return n%2 == 0
}

//reciever function to get even/odd type of number
func (n number) getNumberType() string {
	if n.isEven() {
		return "even"
	} else {
		return "odd"
	}
}

//prints even/odd for given slice of numbers
func printEvenOdd(numbers []number) {
	for _, num := range numbers {
		fmt.Println(num, " is ", num.getNumberType())
	}
}

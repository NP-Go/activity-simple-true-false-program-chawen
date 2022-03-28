package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessge := ""
	secretValue := 66

	//Insert your code from here
	if value == secretValue {
		fmt.Println ("Check 66")
	}

	if value < secretValue {
		fmt.Println ("Low, again")

		
	if value > secretValue {
		fmt.Println ("High, again")

	//do not remove this line
	return resultMessge
}

func main() {
	var guess int
	fmt.Println("Enter integer value (66): ")
	fmt.Scanln(&guess)
	compare(guess)
}

package main

import "fmt"

func main() {
	var feet float64
	str1 := "Enter length (in ft):"
	fmt.Println(str1)
	fmt.Scan(&feet)
	feettometers := feet * 0.3048
	fmt.Println(feet, "feet is", feettometers, "meters.")
}

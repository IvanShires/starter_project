package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var non_alchoholic_drinks [5]string
var max_age int = 86
var min_age int = 21

func non_alchoholic_generate() {
	non_alchoholic_drinks[0] = "Orange Juice"
	non_alchoholic_drinks[1] = "Apple Juice"
	non_alchoholic_drinks[2] = "Cranberry Juice"
	non_alchoholic_drinks[3] = "Water"
	non_alchoholic_drinks[4] = "Gatorade"
}

func main() {
	var name = "Tessa"
	var age int = 12
	non_alchoholic_generate()
	if strings.HasPrefix((name), "Ivan") {
		fmt.Println("Hello,", name)
	} else {
		fmt.Println("Who are you,", name, "?")
	}
	if age < min_age {
		var drink_value int = rand.Intn(3) // Need to figure out how to make this more random
		fmt.Println("You are not allowed to drink", name, ", have some", non_alchoholic_drinks[drink_value], "you can drink in", min_age-age, "year(s)")
	} else if age > max_age {
		fmt.Println(name, "probably not a good idea to drink at your age.. You should've stopped drinking", age-max_age, "year ago!")
	} else {
		fmt.Println("Enjoy the beer,", name, "!")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Namr: ")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Age: ")
	age, _ := reader1.Read("\n")
	fmt.Println(age)

	fmt.Println("Your name is ", name, " and you are age ", age)
}

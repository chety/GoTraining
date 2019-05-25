package main

import "fmt"

func main()  {
	name := true
	var nameType = fmt.Sprintf("%T",name)
	fmt.Println(nameType)
}

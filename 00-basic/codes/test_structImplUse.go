package main

import (
	"Go/00-basic/codes/struct_interface"
	"fmt"
)

func main() {

	name := "Tom"
	fmt.Print("name:", name)

	s, err := struct_interface.New(name)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(s)

	fmt.Println(s.Listen("english"))
	fmt.Println(s.Speak("english"))
	fmt.Println(s.Read("english"))
	fmt.Println(s.Write("english"))
}

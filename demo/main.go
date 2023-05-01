package main

import (
	"fmt"
	"main/mytest"
	"main/mytest/ModuleA"
	ve "github.com/VVVVictorJ/GOPACKAGETEST"
)

func main(){
	fmt.Println("Test")
	fmt.Println(mytest.Test(2, 1))
	fmt.Println(modulea.TestA(2, 1))
	fmt.Println(ve.Version())
}
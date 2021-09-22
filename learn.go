package main

import "fmt"

func main(){
	//fmt.Printf("hello, world")
	stringPrint()
}

func OutSideMethod(){
	fmt.Print("this is an outside accessible method");
}

func stringPrint(){
	var a = "hello"
	var b = "world"
	var c = "haha"
	var result = fmt.Sprintf(a+" %s %s",b,c)
	fmt.Println(result)
}
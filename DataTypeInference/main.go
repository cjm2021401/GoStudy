package main

import "fmt"
//go의 자료형 추론은 c와 마찮가지로 컴파일시점에서 결정
//정적 자료형임 python과 다름
func main(){
	var a int=10
	var a2 =10 
	a3:="little"
	a3+="boy"
	fmt.Println("Hello", a,a2,a3)

}
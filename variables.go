package main

import (
	"errors"
	"fmt")
var var1 int
var var2 uint
var var3, var4 int
var var5 float32
var var6 float64
var var7 complex64
var var8 complex128
var var9 string
var var10 bool
var var11 string = "Hello"
const red = 0
const(
	car = iota
	bus 
	train
)

func main(){
	fmt.Println(car)
	fmt.Println(bus)
	fmt.Println(train) 

	//short decalration
	c:=10
	fmt.Println(c)

	//type conversion
	var3 = 10
	var5=(float32(var3))

	//multivariable
	a,b := 2,3
	fmt.Printf("%v + %v = %v\n",a ,b, a+b)

	//errors
	err := errors.New("Hello")
	fmt.Println(err)

}
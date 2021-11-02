package main

import (
	"fmt"
	"math"
)

func main()  {
	//算数运算符
	num1 := 100
	num2 := 200
	fmt.Printf("num1+num2=%v",num1+num2)
	age := 23
	age++
	fmt.Println(age)
	area := math.Pi * 10 *10
	fmt.Println(area)

	//赋值运算符 =, +=, -=, *=, /=
	var num = 10
	fmt.Println(num)
	num *= 3
	fmt.Println("num的值是：",num)

	//关系运算符
	flag1 := 100 > 200
	flag2 := num == 30
	fmt.Printf("flag1的数据类型：%T; flag2的数据类型：%T\n", flag1, flag2)
	fmt.Println(flag1)
	fmt.Println(flag2)

	//逻辑运算符
	temp1 := 5 >3 || 2 >3
	fmt.Println(temp1)
	var m, n = 10, 20
	temp2 := m >= 10 && n >= 20
	fmt.Println(temp2)





}

package main

import "fmt"

//浮点类型

func main()  {
	//浮点类型【小数类型】
	//float
	var price float32 = 99.99
	fmt.Println(price)
	var a float32 = 123.3
	fmt.Println(a)
	var b float64 = 345.6
	fmt.Println(b)
	
	b = float64(a)//类型转换
	fmt.Println(b)

}

package main

import "fmt"

//值传递：实际是将实参的值复制到形参相应的存储单元
//基本类型数据都是值传递，结构体


//引用传递：也叫地址传递
//切片、指针类型

func main()  {
	a := 10
	b := 20
	fmt.Println(a,b)
	swap(&a, &b)
	fmt.Println(a,b)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

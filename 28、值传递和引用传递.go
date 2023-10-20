package main

import (
	"fmt"
)

//值传递：实际是将实参的值复制到形参相应的存储单元
//基本类型数据都是值传递、结构体、数组


//引用传递：也叫地址传递
//切片、指针类型

type Student1 struct {
	id int
	name string
	score float64
}

func main()  {
	fmt.Println("===基本数据类型===")//值传递
	name := "小明"
	age := 12
	fmt.Println("调用函数前的数据为：", name, age)
	test01(name, age)
	fmt.Println("调用函数后的数据为：", name, age)

	fmt.Println("===数组类型===")//值传递
	arr := [6]int{111,222,333,444,555,666}
	fmt.Println(arr)
	test02(arr)
	fmt.Println(arr)

	fmt.Println("===切片类型===")//引用传递
	s := []int{111,222,333,444,555,666}
	fmt.Println(s)
	test03(s)
	fmt.Println(s)

	fmt.Println("===结构体类型===")//值传递
	student := Student1{1,"小明",12.2}
	fmt.Println(student)
	test04(student)
	fmt.Println(student)

	fmt.Println("===指针类型===")//引用传递
	p := &student
	fmt.Println(p)
	test05(p)
	fmt.Println(p)
}

func test05(p *Student1) {
	p.score = 99.9
	p.name = "阿云嘎"
}

func test04(student1 Student1) {
	student1.id = 2
	student1.name = "小明222"
}

//参数是数组类型
func test03(s []int) {
	s[0] = 888
}

//参数是数组类型
func test02(arr [6]int) {
	arr[0] = 888
}

//参数为基本数据类型
func test01(name string, age int) {
	name = "小刚"
	age = 14

}
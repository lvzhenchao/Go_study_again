package main

import "fmt"

//空接口是指没有定义任何接口方法的接口
//任何类型都实现了空接口，因此空接口可以存储任意类型的数值
//空接口可以理解为万能类型，可以接收任意类型的数据；变量具体的类型依据赋值进行相应的变化
//类似的打印函数：fmt.Println等


func main()  {
	fmt.Println("===空接口的基本使用===")
	var i interface{}
	fmt.Printf("变量i的类型：%T\n", i)
	i = 123
	fmt.Printf("变量i的类型：%T\n", i)
	i = "Go语言"
	fmt.Printf("变量i的类型：%T\n", i)

	fmt.Println("===空接口与切片的应用===")
	s1 := make([]interface{},0)//切片元素可以是任意类型的数据
	s1 = append(s1, "hello", 666, 99.9, true)
	fmt.Println(s1)

	fmt.Println("===空接口与map的应用===")
	map1 := make(map[string]interface{},0)
	map1["001"] = 0x0001
	map1["002"] = "Go语言"
	map1["003"] = s1
	fmt.Println(map1)

	fmt.Println("===空接口作为函数参数===")
	InterfaceTest(123,"go",map1,s1)

}

//可变参数，任意类型
func InterfaceTest(args ...interface{})  {
	for k,v := range args {
		fmt.Printf("第%v个参数，值为：%v，类型为：%T\n", k,v,v)
	}
}
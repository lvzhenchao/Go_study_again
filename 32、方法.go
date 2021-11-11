package main

import "fmt"

//方法和函数的区别：最大的区别有没有接受者

//指针变量的方法值：接收者是结构体的指针类型；节省内存

func test()  {//函数
	fmt.Println("这里是函数")
}

type A1 struct {
	Num int
}

//定义对象（结构体）A1的方法；注意：方法是针对某一对象的函数；只有该类型的对象才能调用
func (a A1)test()  {
	fmt.Println("这里是结构体、对象A1的方法 test()")
	fmt.Println(a.Num)
}

type Person1 struct {
	name string
	sex byte
	age int
}

func (s *Person1)PrintInfo()  {
	fmt.Printf("name:%v;sex:%v;age:%v\n", s.name,s.sex,s.age)
}

func main()  {
	test()

	a := A1{Num:100}
	a.test()

	//正常都是使用这个
	p := Person1{"小明",1,23}
	p.PrintInfo()

}


package main

import "fmt"

//闭包：只有权访问另一个函数作用域中变量的函数；就是在一个函数内部创建另一个函数
//可以理解成：定义在一个函数内部的函数
//闭包是将函数内部函数外部连接的桥梁
//特点：匿名函数会改变外部函数内的变量

func main()  {
	f1 := Test()
	f2 := Test()
	f3 := Test()

	fmt.Println(f1,f2,f3)//1,1,1//每次调用，都是重新分配内存，变量x会重新声明；函数执行完毕，x会自动释放被占用的内存

	fmt.Println("匿名函数执行============")
	f := func()int {
		var i int
		i++
		return i
	}

	fmt.Println(f(), f(), f())//1,1,1
	fmt.Println("闭包函数执行============")

	var x int
	f4 := func() int{
		x += 2
		return x
	}
	fmt.Println(f4(), f4(), f4())//2,4,6
	fmt.Printf("%T", f4)//类型：func() int

	fmt.Println("闭包函数的操作==========")
	f5 := Test1()
	fmt.Println(f5(),f5(),f5())//1,2,3

}

func Test()int  {
	var x int
	x++
	return x
}

func Test1() func()int {
	var x int
	f := func() int{
		x++
		return x
	}
	return f
	//return func() int{
	//		x++
	//		return x
	//	}
}

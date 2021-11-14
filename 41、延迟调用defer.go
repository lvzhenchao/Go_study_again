package main

import "fmt"

//关键字defer用于延迟一个函数【或者当前创建的匿名函数】的执行；栈的形状：先进后出的效果

//注意：defer语句只能出现在函数的内部；代码逻辑越复杂，defer使用越重要

//应用场景：
	//-文件操作：先打开文件，执行读写操作，最后关闭文件
	//-网络编程：最后关闭整个网络的链接

func main()  {
	fmt.Println("===defer的基本使用===")
	defer fmt.Println("Hello Go")
	fmt.Println("开始执行命令")
	defer fmt.Println("执行命令1")
	defer fmt.Println("执行命令2")
	fmt.Println("===defer与匿名函数结合使用===")
	deferTest1()
	deferTest2()

}

func deferTest1()  {
	fmt.Println("开始执行deferTest1")
	a := 10
	b := 20
	defer func() {
		fmt.Printf("匿名函数中a的值为：%v, b的值为：%v\n", a,b)
	}()
	a = 100
	b = 200
	fmt.Println("函数中a，b的值为：", a,b)
}

func deferTest2()  {
	fmt.Println("开始执行deferTest2")
	a := 10
	b := 20
	defer func(a,b int) {
		fmt.Printf("匿名函数中a的值为：%v,b的值为：%v\n", a,b)
	}(a,b)
	a = 100
	b = 200
	fmt.Println("函数中a，b的值为：", a,b)
}



package main

import "fmt"

//多态：通过接口实现；可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态
//多态指的是多种表现形式

//**************
//接口的用法
//	-一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的【任意实现类型对象】作为参数
//	-定义一个类型为接口类型的变量，实际上可以赋值实现类的对象
//**************

type Machine interface {
	run(operate string)(status bool)
}

type Car struct {
	name string
	mode string
}
type Computer struct {
	name string
	cpuCount uint8
}

func (c *Car)run(operate string)(status bool)  {
	fmt.Println(c.name,"car正在执行", operate)
	return true
}

func (c *Computer)run(operate string)(status bool)  {
	fmt.Println(c.name,"computer正在执行", operate)
	return true
}

func main()  {
	car := Car{"长城","越野"}
	run(&car)//实现了多态；注意：传递的变量是car的地址
	run(new(Car))

	computer := Computer{"华为",16}
	run(&computer)

	run(new(Computer))
}

//定义一个普通的函数，函数的参数为接口类型
func run(m Machine)  {//函数实现了多态，注意：形式参数为接口类型
	m.run("启动")////具体执行哪个对象的方法，是由调用函数的对象决定

}



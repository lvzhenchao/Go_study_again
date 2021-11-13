package main

import "fmt"

//多态：通过接口实现；可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态
//多态指的是多种表现形式

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
	computer := Computer{"华为",16}
	run(&computer)
}

//定义一个普通的函数，函数的参数为接口类型
func run(m Machine)  {//函数实现了多态，注意：形式参数为接口类型
	m.run("启动")////具体执行哪个对象的方法，是由调用函数的对象决定

}



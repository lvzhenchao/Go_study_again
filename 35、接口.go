package main

import "fmt"

//接口：就是一种规范与标准

//接口和继承的不同
//	继承的价值：解决代码的复用性和可维护性
//	接口的价值：设计，设计好各种规范（方法），让其它自定义类型去实现这些方法；一定程度上实现代码解耦
//	接口比继承更加灵活

//**************
//接口的用法
//	-一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的【任意实现类型对象】作为参数
//	-定义一个类型为接口类型的变量，实际上可以赋值实现类的对象
//**************

//接口定义
type Humaner interface {
	sayHello() //注意：在接口中的方法只有声明，没有实现；由对象对应的方法来实现
}

type Student3 struct {
	name string
	id int
}

func (s *Student3) sayHello()  {
	fmt.Println("接口的实现")
	fmt.Println(s.name,s.id)
}

//复杂接口
type USB interface {
	Start(operate string)(res byte, err error)//返回值名称可以省略
	Stop() string//单个返回值可以不加括号
}

type Phone struct {
	id int
	name string
}
type Camera struct {
	id int
	mode string
}

func (p *Phone)Start(operate string)(res byte, err error)  {
	fmt.Println(p.name, "phone 正在执行...",operate)
	res = 1
	err = nil
	return
}
func (p *Phone)Stop()string  {
	return "phone停止工作"
}

func (c *Camera)Start(operate string)(res byte, err error)  {
	fmt.Println("编号为",c.id, "的camera正在执行")
	return 1, nil
}
func (c *Camera)Stop()string  {
	return "停止工作"
}

func main()  {
	s := Student3{id:1001, name:"小小"}
	s.sayHello()
	fmt.Println("===复杂接口方法的调用===")
	p := Phone{id:0x01, name:"华为"}
	start, err := p.Start("拍照")//自动添加变量的快捷键ctrl+alt+v
	fmt.Println(start,err)

}

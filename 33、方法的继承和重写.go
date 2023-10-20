package main

import "fmt"

//方法继承：父类的方法可以被子类继承使用
//方法重写：子类的方法名和父类的方法名同名，动用调用的时候是先调用子类的的方法
type PersonA struct {
	name string
	sex byte
	age int
}

func (p *PersonA)Printinfo()  {
	fmt.Println("父类的方法")
	fmt.Println(p.name,p.sex,p.age)
}

type StudentA struct {
	PersonA//匿名字段，主要实现继承
	id int
	addr string
}

func (s *StudentA)Printinfo()  {
	fmt.Println("这是重写的方法")
	fmt.Println(s.name,s.age)
}


func main()  {
	fmt.Println("===方法的继承===")
	s := StudentA{PersonA{"张三", 1, 23}, 20230001, "北京"}
	s.Printinfo()
	fmt.Println("===方法的重写===")
	s.Printinfo()//直接调用重写的方法

	//调用父类的方法,指定父类
	s.PersonA.Printinfo()

}

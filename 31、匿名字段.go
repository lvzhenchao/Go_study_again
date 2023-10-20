package main

import "fmt"

//匿名字段：也称为嵌入字段；只有类型，没有名字
//同名字段：嵌入字段中有一样的成员；原则：先找自己的，没有再找父级的
//指针类型匿名字段：匿名字段带星号*；赋值的时候使用new
//多重继承：多层嵌套；注意：尽量减少多重继承

type person struct {
	id string
	name string
	age int
}
type person1 struct {
	addr string
}

type student2 struct {
	person
	score float64
	name string
	person1
}

func main()  {
	stu := student2{person: person{id:"0102",name:"小刚",age:12},score: 98.9}
	fmt.Printf("变量stu的类型：%T,编号：%v,姓名：%v,年龄：%v,分数：%v\n",stu, stu.person.id,stu.person.name,stu.person.age,stu.score)
	//fmt.Printf("变量stu的类型：%T,编号：%v,姓名：%v,年龄：%v,分数：%v\n",stu, stu.id,stu.name,stu.age,stu.score)
	fmt.Println("===对象属性的操作===")
	//在go语言中，可以将结构体称作类、对象、实体
	stu.score = 88
	stu.age = 23
	stu.age++
	fmt.Println(stu)
	fmt.Println("===子对象和父对象重名===")
	stu.name = "子类名称"
	stu.person.name = "父类名称修改"
	fmt.Println(stu)

	fmt.Println("===多重继承===")//多个父类对象
	stu.person1.addr = "北京"
	fmt.Println(stu)
}


package main

import "fmt"

//指针就是地址【虚拟地址】；
//指针变量：存放地址的变量 new(type)

type student struct {
	id int
	name string
	score float64
}

func main()  {
	var a int = 100
	fmt.Println(&a)

	p := &a
	fmt.Printf("p的变量类型：%T,值为：%v, *P的值为：%v\n",p,p,*p)//p的变量类型：*int,值为：0xc0000ac058, *P的值为：100
	fmt.Println("=====指针变量的含义=====")
	p1 := new(string)
	*p1 = "name"
	fmt.Println(p1, *p1)//0xc000088230 name

	fmt.Println("===指针的使用===")
	str := "Go语言是一门很强大的编程语言"
	p2 := new(string)
	p2 = &str
	fmt.Printf("变量str的值是：%v, *p2的值是：%v\n", str, *p2)

	fmt.Println("===通过指针修改数据===")
	*p2 = "Go语言入门"
	fmt.Printf("*p2的值是：%v, str的值是：%v\n", *p2, str)

	fmt.Println("===结构体指针===")
	p3 := &student{}
	p3.name = "小明"
	p3.id = 1
	p3.score = 40
	fmt.Printf("指针变量p3的值为：%v, *p3的值为：%v\n", p3, *p3)

	p4 :=new(student)
	fmt.Printf("%v", p4)



}

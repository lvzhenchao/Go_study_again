package main

import "fmt"

//指针就是地址【虚拟地址】；
//指针变量：存放地址的变量 new(type)  例如：不同类型的指针 *int *string *bool

//Go语言中的 new 和 make 主要区别如下：
	//make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
	//new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
	//new 分配的空间被清零。make 分配空间后，会进行初始化。

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
	fmt.Printf("%v\n", p4)

	p5 := new(int)
	fmt.Printf("%T,%v",p5,p5)



}

package main

import "fmt"

//结构体：将【不同类型】的数据组合成一个有机的整体
//是一种聚合类型，由一系列具有相同类型或不同类型的数据构成的数据集合
//每个数据称为结构体的成员

//结构体的全部成员都是可以比较的，那么结构体也是可以比较的
//可以使用== 或 != 运算符进行比较  不支持 > 或 <
//同类型的两个机构体变量可以相互赋值

//结构体数组
//用结构体存储多个学生的信息，通过结构体定义多个结构体变量

type Student struct {
	id int
	name string
	sex byte //1:表示男生 0：表示女生
	address string
}

func main()  {
	//初始化结构体
	var s1 Student = Student{1,"阿香",0,"北京"}
	fmt.Printf("s1的类型：%T, 数据为：%v\n", s1, s1)

	//初始化：单独给部分字段赋值
	s2 := Student{name:"小明", sex:1}
	fmt.Println(s2)

	//初始化
	s3 := Student{}
	s3.sex = 0
	s3.name = "小鱼"
	fmt.Println(s3)

	fmt.Println("=====结构体使用=====")

	flag := (s1 == s2)
	fmt.Println(flag)

	fmt.Println("=====结构体数组/切片=====")
	students := make([]Student,0)
	students = append(students, s1,s2,s3)
	fmt.Println(students)

	for _, v := range students{
		fmt.Printf("编号是：%v, 姓名是：%v\n", v.id, v.name)
	}

}

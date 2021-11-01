package main

import "fmt"

//格式化
//%c 字符型。把输入的数字按照ASCLL码转为对应的字符
//%d 一个十进制数值（基数为10）
//%s 字符串。
//%t 输出布尔值
//%T 使用Go语法输出值的类型

//【%v可以代表任意类型】

func main()  {
	name := "小明"
	score := 98
	fmt.Println("姓名："+name)
	fmt.Printf("姓名：%s, 分数：%d", name, score)//格式化输出

	//问题：如果不确定变量的类型怎么办
	fmt.Printf("\n姓名：%v, 分数：%v", name, score)
	fmt.Println("---------------")

	//问题：将变量格式化之后赋值给另外一个变量，如何操作
	str := fmt.Sprintf("姓名：%v,分数：%v", name, score)
	fmt.Println("\nstr变量的内容是："+ str)

	//输入
	var name1 string
	fmt.Println("请输入姓名：")
	fmt.Scan(&name1)//获取输入的内容
	fmt.Println("您输入的姓名是："+name1)

	var age int
	fmt.Scanf("%d", &age)//格式化输入；变量记得写&符号
	fmt.Println("您输入的年龄是：", age)

}

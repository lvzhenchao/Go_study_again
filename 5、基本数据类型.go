package main

import "fmt"

//浮点类型

func main()  {
	//浮点类型【小数类型】
	//float
	var price float32 = 99.99
	fmt.Println(price)
	var a float32 = 123.3
	fmt.Println(a)
	var b float64 = 345.6
	fmt.Println(b)

	b = float64(a)//类型转换
	fmt.Println(b)

	var score1 float32 = 98.9
	var score2 float64 = 89.8
	sum := score1 + float32(score2)
	fmt.Println(score1, score2, sum)//浮点值的加减法 后面会加上不确定位
	fmt.Printf("score1的类型：%T,score2的类型：%T,sum的类型：%T\n", score1, score2, sum)
	fmt.Println("----------")


	//字符类型
	var ch byte	//声明字符类型，关键字byte
	ch = 'a'	//单引号，字符;字符类型用单引号括起来，【只能写一个字符】
	fmt.Printf("%c", ch)
	fmt.Println()
	fmt.Println(ch)//ascll码 97
	fmt.Println('b')//98
	fmt.Println("----------")

	//字符串类型
	//【双引号】括起来的字符串类型
	var str1 string
	str1 = "abc"
	fmt.Println("str = ", str1)
	fmt.Printf("str1=%s", str1)
	fmt.Println()

}

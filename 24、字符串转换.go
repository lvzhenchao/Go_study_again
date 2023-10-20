package main

import (
	"fmt"
	"strconv"
)

//字符串转换包strconv

func main()  {
	//other_to_string()
	//string_to_other()/

	f := 20/3
	fmt.Println(f)//整数/整数 = 整数
	f1 := 20/3.0
	fmt.Println(f1)//6.666666666666667
}

//1、Format系列函数把其他类型的转换为字符串
func other_to_string()  {
	//将bool类型转换为字符串
	var str string
	str = strconv.FormatBool(true)
	fmt.Printf("类型：%T\n", str)
	fmt.Println(str)

	//整形转换为字符串
	var str1 string
	str1 = strconv.Itoa(12)
	fmt.Printf("类型：%T\n", str1)
	fmt.Println(str1)

	var str2 string
	str2 = strconv.FormatFloat(3.14,'f',3,64)//f指打印格式，以小数方式，3指小数点位数，64以float64处理
	fmt.Println(str2)
	fmt.Printf("类型：%T\n", str2)
}

//2、Parse把字符串转换为其他类型
func string_to_other() {
	//转布尔
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(flag)
	} else {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", flag)

	//转整形
	a, _ := strconv.Atoi("567")
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	//转浮点
	b,_ := strconv.ParseFloat("123.456", 64)
	fmt.Println(b)
	fmt.Printf("%T", b)

}

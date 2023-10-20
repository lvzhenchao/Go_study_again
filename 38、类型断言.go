package main

import "fmt"

//判断是否是该类型的变量
//value,ok := element.(T)//element是interface变量，T是断言的类型

func main()  {
	fmt.Println("===类型断言的基本使用===")
	var i interface{} = 123
	var a int
	//var a string
	a,ok := i.(int)
	//a,ok := i.(string)
	if ok {
		fmt.Println("断言成功")
	} else {
		fmt.Println("断言失败")//返回值的为空，相应类型的空值
	}
	fmt.Printf("变量a的值：%v\n",a)

	fmt.Println("===类型断言复杂（常用写法）===")
	if v,ok := i.(int);ok{
		fmt.Println("断言成功！", v)
	} else {
		fmt.Println("断言失败")
	}

	fmt.Println("===类型断言复杂应用===")
	s1 := make([]interface{}, 0)
	s1 = append(s1, 123,98.9,"Go语言", true)
	for k,data := range s1 {
		if v,ok := data.(int);ok {
			fmt.Printf("第%v个元素，数据类型为：%T; 数据为：%v\n", k,v,v)
		} else if v,ok := data.(string);ok {
			fmt.Printf("第%v个元素，数据类型为：%T; 数据为：%v\n", k,v,v)

		} else if v,ok := data.(bool);ok {
			fmt.Printf("第%v个元素，数据类型为：%T; 数据为：%v\n", k,v,v)

		} else if v,ok := data.(float64);ok {
			fmt.Printf("第%v个元素，数据类型为：%T; 数据为：%v\n", k,v,v)

		} else {
			fmt.Println("其他")
		}
	}




}

package main

import "fmt"

//不定参数是指函数传入的参数个数是不确定的数量
//func funcName(v...int)  {
//
//}

//注：不定参数必须放到最后一个参数位置;并且只能有一个不定参数
//func param(arg1 int, arg2...string)



func main()  {
	//方式1：一个一个参数的传递，用逗号隔开
	params1("1","2","3")

	////方法2：将参数设置成切片slice传递
	temp := []string{"4","5", "6"}
	params1(temp...)

	s := make([]string,0)
	s = append(s, "ss")
	params1(s...)//记得加3个小点儿

	fmt.Println("---混合类型不定参数---")
	mixType := make([]interface{},0)
	mixType = append(mixType, "mix1", 23, true,12.3)
	params2(mixType...)
}


func params1(args...string)  {
	for k,v := range args {
		fmt.Println(k,v)
	}
}
func params2(args...interface{})  {
	for k,v := range args {
		fmt.Printf("key:%v, value类型：%\n" ,k,v)
	}
}



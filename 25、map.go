package main

import "fmt"

//map 映射、字典 是一种内置的数据结构
//是一个无序的key-value对的集合，比如以身份证作为唯一键来标识的信息

func main()  {
	//var
	var m = map[int]string{1:"张三", 2:"lisi"}
	fmt.Println(m)

	//自动推导
	m1 := map[int]string{1:"liu", 2:"san"}
	m1[1] = "huan"//修改
	m1[3] = "lzc"//追加
	fmt.Println(m1)

	//make创建
	m2 := make(map[int]string, 10)
	m2[0] = "aaa"
	m2[1] = "bbb"
	fmt.Println(m2)

	//删除delete函数
	delete(m1, 2)
	fmt.Println(m1)

}

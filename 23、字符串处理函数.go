package main

import (
	"fmt"
	"strings"
)

func main()  {
	//str_contains()
	//str_join()
	//str_split()
	//str_fields()
	//str_index()
	//str_replace()
	str_trim()
}

//fields 去除字符串的空格符，并且按照空格分割返回slice
func str_fields()  {
	str := "hello go 语言"
	arr2 := strings.Fields(str)//返回字符串类型的切片
	fmt.Println(arr2)
}

//trim 去除字符串头尾指定的字符
func str_trim(){
	var str = "  go语言是最棒的  "
	fmt.Println(len(str), str)
	trim := strings.Trim(str, " ")
	fmt.Println(trim,len(trim))
}

//分割字符串
func str_split(){
	str := "go语言-java语言-c语言"
	arr1 := strings.Split(str, "-")//返回值是字符串类型的切片
	fmt.Println(arr1)
}

//replace 替换字符
func str_replace(){
	str := "hello go 语言 go"
	replace := strings.Replace(str, "go", "区块链", 2)
	fmt.Println(replace)
}

//repeat 重复字符串几次，返回重复后的字符串
func str_repeat(){
	
}

//index 查找位置，返回位置值,不存在返回-1
func str_index(){
	str := "hello go"
	index := strings.Index(str, "go")
	fmt.Println(index)
}

//join,字符串链接，
func str_join(){
	arr := []string{"go语言","java语言","c语言"}
	str1 := strings.Join(arr, "-")//第一个参数是字符串类型的切片

	fmt.Println(str1)
}

//contains 字符串s中是否包含其他str，返回bool
func str_contains(){
	var str string = "hellogo"
	fmt.Println(strings.Contains(str, "go"))//返回true
	fmt.Println(strings.Contains(str, "abs"))//返回false
}

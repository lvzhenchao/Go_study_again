package main

import "fmt"

func main()  {
	//4种数组初始化

	//最完整、规范的写法
	var arr1[3]int = [3]int{1,2,3}
	fmt.Println(arr1[0])//打印第一个元素
	fmt.Println(arr1[len(arr1)-1])//打印最后一个元素

	//简写版
	var arr2 = [3]int{5,6,7}

	//不确定元素的个数
	var arr3 = [...]int{8,9,10}

	//给具体的某个元素赋值
	var arr4 = [3]int{1:100}

	fmt.Println(arr1,arr2,arr3,arr4)//[1 2 3] [5 6 7] [8 9 10] [0 100 0]

	fmt.Println("--相关补充--")
	arr5 := [...]int{1,2,3}//"..."数组长度的位置出现省略号,则【表示数组的长度】是根据初始化值的个数来计算
	fmt.Println(arr5)
	arr5[2] = 6//不能给超出数组长度的位置赋值，偏移量
	fmt.Println(arr5)





}

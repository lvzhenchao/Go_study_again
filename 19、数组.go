package main

import "fmt"

func main()  {
	//4种数组初始化

	//最完整、规范的写法
	var arr1[3]int = [3]int{1,2,3}

	//简写版
	var arr2 = [3]int{5,6,7}

	//不确定元素的个数
	var arr3 = [...]int{8,9,10}

	//给具体的某个元素赋值
	var arr4 = [3]int{1:100}

	fmt.Println(arr1,arr2,arr3,arr4)//[1 2 3] [5 6 7] [8 9 10] [0 100 0]



}

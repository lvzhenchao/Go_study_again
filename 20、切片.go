package main

import "fmt"

//切片是数组的引用，是引用类型，遵循引用传递的机制
//使用方式和数组类似，遍历，访问切片，求长度
//长度是可变的，是一个可以动态变化的数组
//var 切片名 []类型

//注：make只能创建slice，map和channel，并且返回一个有初始值(非零)的对象

func main()  {
	//切片创建方法
	//1、自动推导类型创建
	s1 := []int{1,2,3}
	fmt.Println(s1)

	//2、借助make创建slice；格式make(切片类型，长度，容量)
	s2 := make([]int, 5,10)
	fmt.Println(s2)

	//3、var
	var s3 []int
	fmt.Println(s3)

	s2[4] = 333

	//s2[6] = 444//panic: runtime error: index out of range [6] with length 5
	//赋值只能是长度范围之内的
	appendSlice()
	copySlice()
}

//切片作为函数参数时，如何传递？传值还是传引用？




//copy在两个切片间复制数据，复制长度以len小的为准，两个切片指向同一底层数组，直接对应【位置覆盖】
func copySlice()  {
	data := [...]int{0,1,2,3,4,5,6,7,8,9}
	s1 := data[8:]//{8,9}
	s2 := data[:5]//{0,1,2,3,4}
	copy(s2, s1)

	fmt.Println(s2)//[8 9 2 3 4]
	fmt.Println(data)//[8 9 2 3 4 5 6 7 8 9]
	fmt.Println("==========")
}

//append()函数可以向切片尾部添加数据，可以自动为切片扩容，【返回新的切片对象】
//切片地址扩容后，切片地址可能发生变化
func appendSlice()  {
	var s1 []int //创建nil切片
	s1 = append(s1, 1)
	s1 = append(s1, 2,3)
	s1 = append(s1, 4,5,6)
	fmt.Println(s1)

	s2 := make([]int, 5)
	s2 = append(s2,6)
	fmt.Println(s2)


	s3 := []int{1,2,3}
	s3 = append(s3, 4,5)
	fmt.Println(s3)

	fmt.Println("==========")
}



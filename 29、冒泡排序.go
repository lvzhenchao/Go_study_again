package main

import (
	"fmt"
	"math/rand"
	"time"
)

//最简单的交换类排序方法，通过相邻数据的交换逐步将线性表变成有序
//基本概念：依次比较相邻的两个数，将小数放在前面，大数放在后面

//每一重循环的目的就是将最大的书移动到最后

func main()  {
	//第一步：先产生随机数
	rand.Seed(time.Now().UnixNano())//通过时间来产生随机数
	arr := make([]int,0)
	for i:=0; i<10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println("切片中的原始数据：", arr)

	BubbleSort(arr)
	fmt.Println("排序之后切片的数据：", arr)

}

func BubbleSort(a[]int)  {//试着画图一下 //5个数的情况下  0轮=>比较4次,1轮=>3次,2轮=>2次,3轮=>1次
	n := len(a)
	//一共有几轮
	for i:=0; i<n-1; i++ {
		//每轮相邻的比较，比较的次数一次比一次少
		for j:=0; j <n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}



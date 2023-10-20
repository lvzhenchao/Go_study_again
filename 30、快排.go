package main

import (
	"fmt"
	"math/rand"
	"time"
)

//快速排序是对冒泡排序的改进
//基本思路：通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分数据要小；然后再按此方法对两部分的数据分别进行快速排序

func main()  {
	//第一步：先产生随机数
	rand.Seed(time.Now().UnixNano())//通过时间来产生随机数
	arr := make([]int,0)
	for i:=0; i<10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println("切片中的原始数据：", arr)

	QuickSort(0,len(arr)-1,arr)
	fmt.Println("排序之后切片的数据：", arr)

}

func QuickSort(left, right int,arr []int) {
	l := left
	r := right
	value := arr[(1+r)/2]//区中间值，作为比较的参考值
	for l <= r {
		for arr[l] < value{//左指针对应的数据满足条件，左指针右移
			l++
		}
		for arr[r] > value{//右指针对应的数据满足条件，右指针左移
			r--
		}
		if l > r {
			break
		}
		arr[l],arr[r] = arr[r],arr[l]//交换数据
		r--
		l++

	}
	if l == r {
		l++
		r--
	}

	if l < right {
		QuickSort(l,right,arr)
	}
	if r > right {
		QuickSort(r,left,arr)
	}
}

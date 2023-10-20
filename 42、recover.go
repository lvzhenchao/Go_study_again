package main

import "fmt"

//运行时panic异常一旦被引发就会导致程序崩溃，专用于【拦截】运行时panic的内建函数--recover
//可以使当前的程序从运行时panic的状态中恢复
//专门捕获panic异常的

//注意：recover只有在【defer】调用的【函数中】有效
//recover一定要写在出现panic的前面

func main()  {
	defer func() {
		v := recover()
		fmt.Println("错误的信息：", v)
	}()//recover调用：1）需要结合defer使用：2）只能在函数内部，一般是匿名函数

	arr := [3]int{}
	i   := 100
	arr[i] = 66
	fmt.Println(arr)//panic: runtime error: index out of range [100] with length 3；数组越界


	fmt.Println("======")//上方如果有panic了，下面还是会停止运行
	recoverTest01(100,0)
	//recover 只能捕获一个panic错误信息，因为：遇到panic错误，程序会结束运行
	//panic错误：如果不用recover捕获是异常退出；捕获之后是正常退出
	//类似php try catch

}

func recoverTest01(a,b float64)float64  {
	if b == 0 {
		panic("被除数不能为0")
		return 0
	}

	return a/b

}

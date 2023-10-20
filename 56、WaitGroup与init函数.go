package main

import (
	"fmt"
	"sync"
)

//WaitGroup主要用于go程的同步
//三个方法：
// Add()：添加协程的数量；
// Done()：相当于Add(-1)；
// Wait()：执行阻塞，直到所有的waitgroup数量变成0

//前期使用for循环

//init函数：每个包完成初始化后自动执行，并且执行优先级比main函数高
//作用：对变量进行初始化、注册、运行一次计算

func init()  {
	fmt.Println("这里是init函数，一般用于初始化操作。。。")
}

func main()  {
	wg := sync.WaitGroup{}//定义
	wg.Add(100)//添加子go程数量
	for i:=0; i<100; i++ {
		go func(n int) {
			fmt.Printf("第%v个go程正在执行操作！\n", n)
			wg.Done()//子go程减一
		}(i)
	}
	//阻塞，直到go程数量 <= 0
	wg.Wait()
}
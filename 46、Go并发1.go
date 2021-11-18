package main

import (
	"fmt"
	"time"
)

//协程比线程小，10几个协程可能体现在底层就是5、6个线程
//特性：主协程退出后，其他子协程也会自动退出

//并发就是执行多个函数，或命令

func main()  {
	fmt.Println("===协程与一般函数的使用===")
	go running()

	for i:=0; i<10; i++ {
		fmt.Println("主进程执行：", i)
		time.Sleep(time.Second*1)
	}

	fmt.Println("===协程和匿名函数的结合===")
	flag := false//标志位
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println("匿名函数执行的并发：", i)
			time.Sleep(time.Second*1)
		}
		flag = true//循环完之后更改值
	}()

	for {
		if flag {
			break
		}
	}
}

func running()  {//如果主程序（主goroutine）退出，并发的goroutine也会退出。【并发程序是寄宿在主程序中的】
	for j:=0; j<100; j++ {
		fmt.Println("子进程执行...", j)
		time.Sleep(time.Second*1)
	}
}

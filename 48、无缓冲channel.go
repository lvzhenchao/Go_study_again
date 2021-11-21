package main

import (
	"fmt"
	"time"
)

//channel是Go语言的核心类型，是一个数据类型
//主要解决go程的同步问题，以及go城之间数据共享（数据传递）问题
//无缓冲通道：指在接收前没有能力保存任何值的通道；
	//发送和接收同时准备好，才能完成发送和接收操作

	//里面只能有一个数据存在，除非被接收了

func main()  {
	c := make(chan int)//定义一个无缓冲的channel;另一种写法：make(chan int, 0)
	fmt.Printf("变量c的类型：%T; 长度为：%v; 变量为：%v\n", c, len(c), cap(c))
	fmt.Println("===channel的基本使用===")

	go func() {
		defer fmt.Println("子go进程退出")
		for i:=0; i<10;  i++{
			c <- i//写入channel，发送
			fmt.Println("变量c的类型：%T; 长度为：%v; 变量为：%v\n", c, len(c), cap(c))
			time.Sleep(time.Microsecond*300)
		}
	}()

	//主程序接收channl
	for i:=0; i<10; i++ {//发送了10次，也得接收10次；如果接收处次数多，会造成死锁；
		num := <- c
		fmt.Println("num=", num)
		time.Sleep(time.Microsecond*300)
	}
}

package main

import (
	"fmt"
	"time"
)

//带缓冲的通道：并不强制要求必须同时完成发送和接收


//有缓冲：在接收前能存储一个或多个数据值的通道；这种类型【不强制必须】同时完成接收和发送
	//通道阻塞发送和接收的条件不同
	//只有通道中没有要接收的值时，接收动作才会阻塞
	//只有通道没有可用的缓冲区【达到数量】容纳被发送的值时，发送动作才会阻塞

//如果给定了一个缓冲区容量，通道就是异步的；只要有未使用的空间，其通信就会无阻塞的进行
//len(ch)求取缓冲区中剩余元素的个数；cap(ch)求取缓冲区容量的大小
//关闭channel：发送者知道没有值需要发放到channel。那么让接受者知道没有多余的值可接收是有用的；
//因此接受者可以停止不必要的接收等待；通过内置的close函数来关闭channel实现
	//关闭后，无法向channel再发送数据
	//关闭后，可以继续从channel接收数据
	//对于nil channel，无论收发都会阻塞

	//len(ch)求取缓冲区剩余元素的个数；cap(ch)求取缓冲区元素容量大小

var ch chan int
var flag = false

func main()  {
	fmt.Println("===有缓冲channel的基本使用===")
	c := make(chan int, 3)
	fmt.Printf("创建有缓冲channel长度为：%v;容量为：%v\n", len(c),cap(c))

	//go程向channel写入数据
	go func() {
		for i:=0; i<10; i++{
			c<-i
			fmt.Printf("通道长度为：%v;容量为：%v\n", len(c), cap(c))
		}
	}()

	//在主程序中读取channel中的数据
	for i:=0; i<10; i++ {
		time.Sleep(time.Second*1)
		num :=<-c
		fmt.Println("num=", num)
	}

	fmt.Println("===有缓冲channel的关闭与遍历===")
	ch = make(chan int, 3)
	go sendData()
	//go recvData1()
	go recvData2()
	for{
		if flag {
			break
		}
	}

}

func sendData()  {
	for i:=0; i<5; i++ {
		time.Sleep(time.Second*1)
		ch <- i*i
	}
	close(ch)
}

func recvData1()  {
	for{
		if data,ok := <-ch; ok {
			fmt.Println("接收到的数据为：", data)
		} else {
			fmt.Println("通道已关闭！")
			break
		}
	}
	flag = true
}

func recvData2()  {
	for data := range ch{//for-range 会一直阻塞读取channel中的数据，直到channel关闭才会退出
		fmt.Println("读取到数据为：", data)
	}
	fmt.Println("子go程recvData2()执行结束")
	flag = true
}
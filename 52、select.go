package main

import (
	"fmt"
	"time"
)

//关键字select，监听channel上的数据流动
//限制：最大限制：每个case语句里必须是一个IO操作
//从头至尾评估每一个发送和接收的语句

//基本应用：超时

func main()  {
	//ch := make(chan int, 1)
	//for {
	//	select {
	//	case ch <- 0:
	//	case ch <- 1:
	//	}
	//	i := <-ch
	//	fmt.Println("Value received:", i)
	//}

	fmt.Println("===select的基本使用===")
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second*3)
		c1 <-"one"
	}()
	go func() {
		time.Sleep(time.Second*2)
		c2<-"two"
	}()
	for i:=0; i<2; i++ {
		select {
			case msg:=<-c1:
				fmt.Println("接收到的数据为：", msg)
			case msg1:=<-c2:
				fmt.Println("接收到的数据为：",msg1)
		}
		fmt.Println("**********")
	}

	fmt.Println("===超时操作===")
	TimeOutOperation()


}

func TimeOutOperation()  {
	c := make(chan int)
	out := make(chan bool)
	go func() {
		for{
			select {
			case v:=<-c:
				fmt.Println(v)
			case <-time.After(time.Second*5):
				fmt.Println("超时！")
				out<-true//标志位
				return//退出

			}
		}
	}()

	//主进程
	<-out//一直等待状态，标志位为阻塞状态
}

















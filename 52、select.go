package main

import "fmt"

//关键字select，监听channel上的数据流动
//限制：最大限制：每个case语句里必须是一个IO操作
//从头至尾评估每一个发送和接收的语句

//基本应用：超时

func main()  {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
}



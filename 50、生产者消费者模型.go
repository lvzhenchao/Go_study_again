package main

import (
	"fmt"
	"time"
)

//单向channel，只能让它发送数据，或只能让它接收数据；可以指定通道的方向

//单向channel最典型的应用是："生产者消费模型"；
//	生产者==>缓冲区==>消费者
//	缓冲区作用：解耦、处理并发、缓存


var flag1 = false

func main()  {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	for {
		if flag1 {
			break
		}
	}
}

//生产者
func producer(ch chan<- int)  {
	for i:=0; i<10; i++ {
		fmt.Printf("生产者第%v条数据\n", i+1)
		ch <- i+1
		time.Sleep(time.Second*1)
	}
	close(ch)
}

//消费者
func consumer(ch <-chan int)  {
	for num := range ch {
		fmt.Println("消费者读取到的内容为：", num)
	}

	flag1 = true
}
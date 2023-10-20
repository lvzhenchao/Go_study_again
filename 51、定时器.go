package main

import (
	"fmt"
	"time"
)

//time.Timer   time.Ticker

//Timer是一个定时器，代表未来的一个单一事件，告诉Timer要等待多长时间
//type Timer struct {
//	C <-chan Time
//	r runtimeTimer
//}
//提供一个channel，在定时时间到达之前，没有数据写入timer.C会一直阻塞，直到定时时间到，【系统】会自动向 timer.C这个channel
//中写入当前时间，阻塞即被解除

//Ticker是一个周期触发定时的计时器，会按照一个时间间隔往channel发送系统当前时间，而channel的接收者可以固定的时间间隔从channel中
//读取时间
//type Ticker struct {
	//C <-chan Time
	//r runtimeTimer
//}

var flag2 = false

func main()  {
	fmt.Println("===定时器的基本使用===")
	timer := time.NewTimer(time.Second*5)
	t1 := time.Now().Second()
	fmt.Println("当前时间是：", t1)
	timer.Reset(time.Second*1)
	t2 := <-timer.C//读取内容
	defer timer.Stop()
	fmt.Println("时间到！", t2.Second())

	fmt.Println("===3种延时操作===")
	delay01()
	delay02()
	delay03()

	fmt.Println("===ticker的使用===")
	ticker01 := time.NewTicker(time.Second*1)//完成周期性的延时操作
	i := 0
	go func() {
		for{
			<-ticker01.C
			i++
			fmt.Printf("执行第%v次\n", i)
			if i>=5 {
				ticker01.Stop()
				flag2 = true
				break
				//return
				//runtime.Goexit
			}
		}
	}()

	for {
		if flag2 {
			break
		}
	}
}

func delay01()  {//复杂的可能需要延时、重置、停止等
	fmt.Println("开始执行延时函数1")
	t := time.NewTimer(time.Second*3)
	<-t.C
	fmt.Println("延时函数1执行结束")
}

func delay02()  {
	fmt.Println("开始执行延时函数2")
	<-time.After(time.Second*3)
	fmt.Println("延时函数2执行结束")

}

func delay03()  {//常用简单的
	fmt.Println("开始执行延时函数")
	time.Sleep(time.Second*3)
	fmt.Println("延时函数3执行结束")
}


































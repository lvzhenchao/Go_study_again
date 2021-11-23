package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//条件变量：类似多个生产者和多个消费者，共享资源【channel】的使用
 	//1、共享资源是空的时候：【通知】阻塞在某个条件上的go程【生产者或消费者】
 	//2、共享资源不是空的：【通知】阻塞在某个条件上的go程【生产者或消费者】

 	//sync.Cond类型代表了条件变量
 	//条件变量要与锁（互斥锁，或读写锁）一起使用

//条件变量对应的3个常用方法
// Wait()：阻塞等待条件变量满足
// Signal()：单发通知，给正在等待【阻塞】在该条件上的协程发送通知
// Broadcast()：广播通知，给正在等待【阻塞】在该变量上所有的协程发送通知


var cond sync.Cond//定义条件变量

func main()  {
	cond.L = new(sync.Mutex)//初始化条件变量的锁
	buf := make(chan int, 3)//生产者消费者模型的缓冲区

	rand.Seed(time.Now().UnixNano())//初始化随机种子

	for i:=0; i<3; i++ {
		go producer01(buf, i)
	}
	for i:=0; i<3; i++ {
		go consumer01(buf, i)
	}

	//主程序
	for {
		time.Sleep(time.Second*1)
		runtime.GC()//垃圾回收
	}
}

//生产者
func producer01(ch chan<- int, n int){
	for{
		cond.L.Lock()//加锁
		for len(ch) == 3 {//缓冲区满了，需等待
			cond.Wait()
		}
		num := rand.Intn(100)
		ch <- num
		fmt.Printf("第%v个生产者正在消费，生产的数据为：%v\n", n, num)
		cond.L.Unlock()//解锁
		cond.Signal()//发送消息至消费者
		time.Sleep(time.Second*2)
	}
}

//消费者
func consumer01(ch <-chan int, n int)  {
	for{
		cond.L.Lock()
		for len(ch) == 0 {//缓冲区空了，需等待
			cond.Wait()
		}
		num := <-ch
		fmt.Printf("第%v个消费者正在消费，消费的数据为：%v\n", n, num)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Second*1)
	}
}
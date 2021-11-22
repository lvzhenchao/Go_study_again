package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁可以让多个读操作并发，同时读取，但是对于写操作完全互斥

//当一个协程进行写操作的时候，其他协程既不能读操作，也不能写操作

//读锁 可以多个
//Rlock && RUnlock

//写锁：只能有一个
//Lock && Unlock

var count = 0//读写锁操作的数据资源
var rwlock sync.RWMutex

func main()  {
	for i:=0; i<5; i++ {//开启5个go程执行读操作
		go read(i+1)
	}
	for i:=0; i<5; i++ {//5个go程写操作
		go write(i+1, i*i)
	}

	select {
	case <-time.After(time.Second*30):
		fmt.Println("主程序（go程）退出")
	}
}

//读操作（读取数据资源）
func read(n int)  {
	rwlock.RLock() //加锁：读模式锁
	defer rwlock.RUnlock()//解锁
	data := count
	fmt.Printf("------读操作正在执行第%v个go程的操作，读取到的数据为：%v\n", n, data)
	time.Sleep(time.Second*3)
}

//写操作（修改数据资源）
func write(n, data int) {
	rwlock.Lock()
	defer rwlock.Unlock()
	count = data
	fmt.Printf("写操作正在执行第%v个go程的操作，写入的数据为：%v\n", n, data)
	time.Sleep(time.Second*3)
}

//读写锁可以并发的读，但是同时只能有一个go程写：并且读、写不能同时进行


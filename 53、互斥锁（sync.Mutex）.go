package main

import (
	"fmt"
	"sync"
	"time"
)

//锁就是某个协程在访问某个资源时先锁住，防止其他协程的访问，等待完毕解锁后其他协程再来加锁进行访问

//互斥锁：是传统并发编程对共享资源进行访问控制的主要手段
//两个公开的指针方法：Lock和Unlock
//借助defer,锁定后，立即使用defer语句保证互斥锁即时解锁

var f1 = false
var f2 = false

var mutex  sync.Mutex

func main()  {
	go person01("hello")
	go person02("world")
	//未加锁，此时会乱序

	for{
		if f1 == true && f2 == true {
			break
		}
	}

}


func printer(str string)  {
	mutex.Lock()//加锁
	defer mutex.Unlock()//解锁
	for _,data := range str{
		fmt.Printf("%c", data)
		time.Sleep(time.Second*1)
	}
	fmt.Println()
}
func person01(str string)  {
	printer(str)
	f1 = true
}
func person02(str string)  {
	printer(str)
	f2 = true
}
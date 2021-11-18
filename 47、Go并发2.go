package main

import (
	"fmt"
	"runtime"
)

//runtime包3个函数
	//runtime.Gosched：让出时间片，让出当前协程的执行权限
	//多条并发执行的时候，1，2，3 ；1不执行，先让2,3执行
	//一条并发，这个函数不起作用

	//runtime.Goexit：立即终止当前协程，

	//runtime.GOMAXPROCS：用来设置可以并行计算CPU核数的最大值

func main()  {
	fmt.Println("计算机的CPU个数：", runtime.NumCPU())
	runtime.GOMAXPROCS(8)//计算机底层会默认1~8的范围，不管超出这个范围
	go func() {
		for i:=0; i<2; i++{
			fmt.Println("hello")
		}
	}()

	go running02()
	go running03()

	for i:=0; i<2; i++ {
		runtime.Gosched()//让出当前的cpu，让该cpu执行其他的go程
		fmt.Println("world")
	}


}

func running02()  {
	for i:=0; i<10; i++ {
		fmt.Println("执行的Go程1", i)
		if i > 6 {//退出
			return
		}
	}
}

func running03()  {
	for i:=0; i<10; i++ {
		fmt.Println("执行的Go程2", i)
		if i > 3 {//退出
			runtime.Goexit()
		}
	}
}



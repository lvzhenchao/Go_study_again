package main

import (
	"fmt"
	"sync"
)

//不断对map进行读和写，会出现错误。主要原因是对map进行读和写发生了竞态问题
//加锁机制、channel同步机制；性能不高

//sync.map是一种效率较高的并发安全机制

//同时操作map，这个时候map有安全机制，发现多个go程操作它，就会报错 fatal error: concurrent map writes
//总结：map在并发条件下，多go程是不安全的

var m map[int]int //定义了一个基本map
var wg sync.WaitGroup//用于go程之间的同步

var RWLock sync.RWMutex

var m1 sync.Map//定义一个sync.map

func main()  {
	m = make(map[int]int)//初始化map

	//wg.Add(10)
	////开辟go程  是有错误的
	//for i:=0; i<10; i++ {
	//	go w1(i, i*i)
	//}

	//第二种方法
	//wg.Add(20)
	//for i:=0; i<10; i++ {
	//	go w2(i,i*i)
	//}
	//for i:=0; i<10; i++ {
	//	go r2(i)
	//}

	//第三种方法
	wg.Add(20)
	for i:=0; i<10; i++ {
		go w3(i,i*i)
	}
	for i:=0; i<10; i++ {
		go r3(i)
	}

	wg.Wait()
}

func w1(key, value int)  {
	m[key] = value
	wg.Done()
}

func w2(key, value int) {
	RWLock.Lock()
	defer RWLock.Unlock()

	m[key] = value
	wg.Done()
}

func r2(key int) {
	RWLock.RLock()
	defer RWLock.RUnlock()
	data := m[key]
	fmt.Printf("map读取到的数据为：%v\n", data)
	wg.Done()
}

func w3(key, value int) {
	m1.Store(key,value)//写数据
	wg.Done()
}

func r3(key int) {
	value, _ := m1.Load(key)//读数据
	fmt.Printf("sync.map读取到的数据为：%v\n", value)
	wg.Done()
}

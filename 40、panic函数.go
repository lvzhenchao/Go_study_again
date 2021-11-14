package main

import (
	"errors"
	"fmt"
)

//error返回的是一般性的错误
//panic函数返回的是让程序崩溃的错误，致命错误
	//-当遇到不可恢复的错误状态的时候，如数组访问越界、空指针引用；这些错误会引起panic错误
	//-当某些不应该发生的场景发生时，应该调用panic
//实际开发过程中并不会调用panic函数，但是在当我们编程的程序遇到致命错误时，系统会自动调用该函数，系统内置了panic函数

//特别注意：一旦调用panic函数，程序会终止，不会继续执行

func main()  {
	panic1 := testPanic1()
	fmt.Println(panic1)
	//testPanic2()

	f := dev1(2, 0)
	fmt.Println(f)
}

func dev1(a,b float64)float64  {
	if b == 0 {
		panic("被除数不能为0")
	}
	return a/b
}

func testPanic1()error  {
	fmt.Println("error接口的函数")
	return errors.New("错误类型为error接口的函数")
}

func testPanic2(){
	fmt.Println("执行panic函数")
	panic("程序发生了严重的错误！")
}
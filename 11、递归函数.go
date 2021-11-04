package main

import "fmt"

//当一个函数在其函数体调用自身，则称之为递归

func main()  {

	fmt.Println(Fac(4))

	fmt.Println(fbn(6))

	fmt.Println(sum(100))
}

//阶乘
func Fac(n int) int {
	if n >0 {
		return n*Fac(n-1)
	}
	return 1
}

//斐波那契数列
func fbn(n int) int  {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func sum(n int) int  {
	if n == 1 {
		return 1
	} else {
		return n+sum(n-1)
	}
}

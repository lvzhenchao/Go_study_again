package main

import (
	"errors"
	"fmt"
	"strconv"
)

//错误处理的标准模式，即error接口，是go语言内建的接口类型
//定义如下：
//type error interface {
//	Error() string
//}

//func New(text string) error {
//	return &errorString{text}
//}
//
//type errorString struct {
//	s string
//}
//
//func (e *errorString) Error() string {
//	return e.s
//}

func main()  {
	fmt.Println("===基本使用===")
	err := errors.New("这是一个错误信息")//实例化
	fmt.Println(err.Error())//调用实话化后的方法
	fmt.Println(err)//调用实话化后的方法

	str := "go语言error的接口"
	e := fmt.Errorf("这里是错误信息，对应的错误为：%v", str)
	fmt.Println(e.Error())

	fmt.Println("===函数调用使用===")

	res, err := Dev(200, 0)
	if err == nil {
		fmt.Println(res)
		return
	}
	fmt.Println(err.Error())
	fmt.Println(err)

}

func Dev(a, b float64)(float64, error)  {
	if b == 0 {
		return 0, errors.New("被除数不能为0")
	}

	str := fmt.Sprintf("%.3f", a/b)
	return strconv.ParseFloat(str, 64)
}
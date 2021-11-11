package main

import "fmt"

//var dog Dog
//掉用  dog.Bark()

//方法值调用： var dog Dog
//dfunc := dog.Bark //保存方法地址
//dfunc()

type operate struct {
	num1 int
	num2 int
}

func (o operate)Add()int   {
	return o.num1 + o.num2
}
func (o *operate)Sub()int  {
	return o.num1-o.num2
}
func main()  {

	fmt.Println("===一般方法的调用====")
	o:= operate{}
	o.num1 = 456
	o.num2 = 123
	fmt.Println(o.Add(), o.Sub())

	fmt.Println("===值方法的调用===")
	f1 := o.Add
	f2 := o.Sub
	fmt.Println(f1,f2,f1(),f2())//0x6d7c40 0x6d7c80 579 333

	fmt.Println("===方法表达式的调用===")
	f3 := operate.Add
	res1 := f3(o)//传值

	f4 := (*operate).Sub
	res2 := f4(&o)//传值
	fmt.Println(res1, res2)
}
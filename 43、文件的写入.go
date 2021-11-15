package main

import (
	"fmt"
	"io"
	"os"
)

//步骤：
//1、导入os包，创建文件，读写文件的函数都在改包
//2、指定创建的文件存放路径及文件名
//3、执行create()函数，进行文件创建
//4、关闭文件

//writeString()//文本文件
//write()//二进制文件；图片，视频，音频
//writeAt//指定位置写入数据

func main()  {
	path := "E:/GoPath/src/Go_study_again/go.txt"
	//path := "E:\\GoPath\\src\\Go_study_again"
	file1 := CreateFile1(path)
	fmt.Println(file1)
}

func CreateFile1(path string) error  {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("文件创建失败")
		return err
	}
	defer file.Close()
	for i:=0; i < 10; i++ {
		file.WriteString("Go 语言是一门很强大的语言\n")
	}

	file.WriteAt([]byte("文件操作"), 0)//添加位置一定注意，不能随意添加；对于文本，加在头或尾

	n, err := file.Seek(0, io.SeekEnd)//移动光标到最后
	if err != nil {
		return err
	}
	file.WriteAt([]byte("123456"),n)
	return nil
}

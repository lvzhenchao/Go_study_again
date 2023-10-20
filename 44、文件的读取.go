package main

import (
	"fmt"
	"io"
	"os"
)

//读取文件的基本流程
//1、打开要读取的文件
//2、对文件进行读取
//3、关闭文件

//open()打开,只有读的权限
//openFile()，可以对文件操作，例如修改

func main()  {
	path := "E:/GoPath/src/Go_study_again/go.txt"
	//ReadFile1(path)
	ReadFile2(path)
}

func ReadFile1(path string)error  {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败")
		return err
	}
	defer file.Close()

	//创建缓冲区
	buf := make([]byte, 1024*4)//创建4k大小的切片缓冲区

	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("文件读取错误")
		return nil
	}
	fmt.Println("读取到内容：", string(buf[:n]))
	return nil
}

func ReadFile2(path string)error  {
	file, err := os.OpenFile(path,os.O_RDWR, 666)
	if err != nil {
		fmt.Println("文件打开失败")
		return nil
	}
	defer file.Close()

	buf := make([]byte, 1024)
	//通过分片处理大文件
	for {
		n,err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("文件读取失败:", err)
			return nil
		}
		if n == 0 {//读到文件末尾了
			break
		}
		fmt.Println(string(buf[:n]))
	}
	return nil

}

package main

import (
	"encoding/json"
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
	//path := "E:\\GoPath\\src\\Go_study_again"
	// 这里不能写成 b := []byte{"Golang"}，这里是利用类型转换。
	//b := []byte("Golang")
	//fmt.Printf("b类型：%T\n", b)

	//这个是类型转换，将字符串“ABC”转为[]byte类型
	//c := []byte("ABC€")
	//fmt.Println(c) // [65 66 67 226 130 172]
	//s := string([]byte{65, 66, 67, 226, 130, 172})
	//fmt.Println(s) // ABC€
	//
	//d := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//fmt.Printf("d的类型：%T\n", d)
	//fmt.Println(d)


	path := "E:/GoPath/src/Go_study_again/go.txt"
	file1 := CreateFile1(path)
	fmt.Println(file1)

	jsonFile := CreateJsonFile()
	fmt.Println(jsonFile)

	ReadJson()
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

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

//生成json格式的文件
func CreateJsonFile() error {
	info := []Website{{
		"Golang", "http://c.biancheng.net/golang/",
		[]string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}},
		{"Java", "http://c.biancheng.net/java/",
			[]string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}

	file, e := os.Create("E:/GoPath/src/Go_study_again/info.json")
	if e != nil {
		fmt.Println("文件创建失败")
		return e
	}
	defer file.Close()

	//创建json编码器
	encoder := json.NewEncoder(file)

	//编码
	e = encoder.Encode(info)
	if e != nil {
		fmt.Println("编码错误：",e)
	} else {
		fmt.Println("编码成功")
	}

	return e
}

func ReadJson() error {
	file, e := os.Open("E:/GoPath/src/Go_study_again/info.json")
	if e != nil {
		fmt.Println("文件打开失败")
		return e
	}
	defer file.Close()

	//创建切片缓冲区
	var info []Website
	//创建解码器
	decoder := json.NewDecoder(file)
	e = decoder.Decode(&info)//存入缓冲区

	if e != nil {
		fmt.Println("解码失败")
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}

	return e

}
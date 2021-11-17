package main

import (
	"fmt"
	"io"
	"os"
)

//实现步骤：
//1、打开源文件
//2、创建新文件
//3、创建缓冲区
//4、多次执行：源文件->缓冲区->新文件

//注：操作完成后，关闭文件
func main()  {
	var srcPath string
	var dstPath string
	fmt.Println("请输入源文件的名称：")
	fmt.Scan(&srcPath)
	fmt.Println("请输入目标文件的名称")
	fmt.Scan(&dstPath)
	fmt.Println(srcPath, dstPath)
	if srcPath == dstPath {
		fmt.Println("路径不能相同")
		return
	}
	copyFile(srcPath, dstPath)
//	F:/百度网盘下载内容/WWW.zip
//	E:/GoPath/src/WWW.zip
}

func copyFile(srcPath, dstPath string) error {
	//打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("打开文件错误！", err)
		return err
	}
	defer srcFile.Close()

	//创建新文件
	dstFile, err := os.Create(dstPath)
	if err != nil {
		fmt.Println("创建文件失败")
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, 4*1024)
	for {
		//读取文件, 将源文件读取到缓冲区
		Read, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("读取源文件错误", err)
			return err
		}
		if Read == 0 {
			fmt.Println("文件读取完毕")
			break
		}

		//将缓冲区中的数据写入目标文件
		dstFile.Write(buf[:Read])


	}
	return nil



}
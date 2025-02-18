package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("D:/AccessOK/code/golang/src/gocode/basic/10fileManage/test.txt")
	if err != nil {
		fmt.Println("文件打开失败！\n")
	}
	fmt.Printf("文件=%v", file)
	//操作完毕之后关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("文件关闭失败")
	} else {
		fmt.Println("文件关闭成功")
	}
}

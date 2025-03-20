package main

import (
	// _ 是对包做初始化操作，但是并不适用包里的标识符
	//Go编译器不允许声明却不使用，下划线让编译器接受这类导入
	//并且调用包内所有代码文件里定义的init函数。
	_ "accessok/inAction/goinaction/code/chapter2/sample/matchers"
	"accessok/inAction/goinaction/code/chapter2/sample/search"
	"log"
	"os"
)

// init在main之前调用
func init() {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main是整个程序的入口

func main() {
	//匹配器：包含特定信息，用于护理某类数据源的实例。
	search.Run("president")
}

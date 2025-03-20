package search

import (
	"log"
	"regexp"
	"sync"
)

/*
注册用于搜索的匹配器的映射
没有定义在任何函数作用域内，会被当成包级变量
声明为Mather类型的映射，这个映射一string类型作为键，Mather作为值
*/

var matchers = make(map[string]Matcher)

func Run(searhTerm string) {
	//获取需要所有的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	//创建一个无缓冲的通道，接受匹配后的结果
	results := make(chan *Result)
	//构造一个waitGroup,以便处理所有的数据源
	var waitGroup sync.WaitGroup

}

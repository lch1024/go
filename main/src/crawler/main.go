package main

import (
	"crawler/fetcher"
	"fmt"
	"regexp"
)

// 构架爬虫 (爬相亲网 http://www.zhenai.com/zhenghun)
//                          城市列表           -----> 根据网页主业html解析城市列表正则表达式解析器
//                         /        \
//                      城市1       城市2      -----> 城市正则表达式解析器
//                     /     \     /     \
//                  用户   用户  用户   用户   -----> 用户正则表达式解析器

// 1.获取主页html
// 2.获取城市信息(名称和链接)
// 方法：使用css选择器
//       使用xpath库
//       使用正则表达式
// 本项目使用正则表达式
// 3.使用正则表达式获取到城市名称 和 链接
// 4.抽象解析器(1.城市列表解析器 2.城市解析器 3.用户解析器)
// 解析器：1.输入 utf-8编码文本 2.输出 Request{URL, 对应Parser}列表, Item列表

func main() {
	all, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if nil != err {
		panic(err)
	}
	printCityList(all)

}

func printCityList(contents []byte) {
	// 正则表达式下面的东西
	//<a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f>阿坝</a>
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
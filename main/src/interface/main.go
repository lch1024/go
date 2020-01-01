package main

import (
	"fmt"
	"interface/mock"
	"interface/real"
	"time"
)

// 接口的定义

/*    使用者      ----------->     实现者
     download					 retriever

	1.接口由 使用者 定义
    2.接口的实现是隐式的
	3.接口变量自带指针
	4.接口变量同样采用值传递，几乎不需要使用接口的指针
	5.指针接收者实现智能以指针方式使用;值接收者都可

	查看接口变量
	6.表示任何类型: interface{}
	7.Type Assertion
	8.Type Switch

	特殊接口(系统提供的 我们可以实现对应接口 达到多态目的)
	9.Stringer
	10.Reader/Writer
 */

// 接口声明 定义实现者必须要实现的函数
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.github.com")
}

func post(p Poster) {
	p.Post("http://www.github/com", map[string]string{
		"name": "lch",
		"course": "golong",
	})
}

// 组合
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.github.com"
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked github.com",
	})

	return s.Get(url)
}

// 常用接口

func main() {
	var r Retriever
	mockRetriever := mock.Retriever{"this is fake github.com"}
	r = &mockRetriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "lch",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
//	fmt.Println(download(r))

	fmt.Println("Try a session")

	fmt.Println(session(&mockRetriever))

}

func inspect (r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("Timeout:", v.TimeOut)
	}
}

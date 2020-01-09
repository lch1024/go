package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// http
// 使用http客户端发送请求
// 使用http.Client控制请求头部
// 使用httputil简化工作

// http服务器的性能分析
// import _ "net/http/pprof"
// 访问/debug/pprof/
// 使用go tool pprof分析性能
// net/http/pprof 看系统库这个文件的 注释有各种新能测试的 cmd说明

// 其他标准库
// bufio
// log
// encoding/json
// regexp
// time
// strings/math/rand
// godoc -http :8888


func main() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	// 验证重定向
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	//resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get("http://www.imooc.com")
	defer resp.Body.Close()
	if nil != err {
		panic(err)
	}

	s, err := httputil.DumpResponse(resp, true)
	if nil != err {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}

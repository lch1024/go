package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// 接口的实现者
type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

// 接口函数的实现
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if nil != err {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()

	if nil != err {
		panic(err)
	}

	return string(result)
}

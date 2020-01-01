package mock

import "fmt"

type Retriever struct {
	Contents string
}

// 如果不传指针 那么设置会失败
func (r *Retriever) Post(url string, resp map[string]string) string {
	r.Contents = resp["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

// 实现fmt/print.go 内接口
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

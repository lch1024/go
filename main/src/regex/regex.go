package main

import (
	"fmt"
	"regexp"
)

const text = `My email is 360225013@qq.com 
			  My email is lnyklblch@126.com 
			  My email is lnyklblch@gmail.com
			  My email is lnyklblch@gmail.com.cn`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}

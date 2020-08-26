package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	// 正则表达式下面的东西
	//<a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f>阿坝</a>
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Request = append(result.Request, engine.Request{Url: string(m[1]), ParserFunc: engine.NilParser})
	}

	return result
}

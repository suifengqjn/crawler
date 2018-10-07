package parser

import (
	"fmt"
	"goDemo/Project/crawler/singleCrawler/engine"
	"regexp"
)

const cityList  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParserCityList(contents []byte) engine.ParserResult  {

	re := regexp.MustCompile(cityList)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _,m := range matches {

		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			engine.NilParser,
		})
		fmt.Printf("City:%s,URL:%s\n", m[2], m[1])
	}
	return result
}

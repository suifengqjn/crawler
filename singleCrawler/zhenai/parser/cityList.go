package parser

import (
	"fmt"
	"goDemo/Project/crawler/singleCrawler/engine"
	"regexp"
)

var CityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
func ParserCityList(contents []byte) engine.ParserResult  {

	matches := CityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	limit := 20
	for _,m := range matches {

		fmt.Println("fetching url", string(m[1]))
		result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			ParseCityUserList,
		})
		limit --
		if limit ==0 {
			break
		}
	}
	return result
}

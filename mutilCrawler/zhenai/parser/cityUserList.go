package parser

import (
	"fmt"
	"goDemo/Project/crawler/singleCrawler/engine"
	"regexp"
)

var cityUserListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

// 获取城市里面的用户列表
func ParseCityUserList(contents []byte) engine.ParserResult {
	matches := cityUserListRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items,"User " + string(m[2]))
		fmt.Println("url", string(m[1]),"name:",string(m[2]))
		name := string(m[2])
		result.Requests = append(result.Requests,engine.Request{
			string(m[1]),
			func(bytes []byte) engine.ParserResult{
				return ParseProfile(bytes, name)
			},
		})
	}

	//// 获取用户列表页面的城市
	//matches = CityListRe.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	result.Requests = append(
	//		result.Requests,
	//		engine.Request{
	//			Url:       string(m[1]),
	//			ParserFunc:ParserCityList,
	//		})
	//}

	return result
}

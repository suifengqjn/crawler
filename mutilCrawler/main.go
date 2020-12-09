package main

import (
	"goDemo/Project/crawler/singleCrawler/config"
	"goDemo/Project/crawler/singleCrawler/engine"
	"goDemo/Project/crawler/singleCrawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		config.CRAWLER_URL,
		parser.ParserCityList,
	})

}

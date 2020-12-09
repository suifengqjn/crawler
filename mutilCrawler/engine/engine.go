package engine

import (
	"fmt"
	"goDemo/Project/crawler/singleCrawler/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			fmt.Println(item)
		}

	}
}

func worker(r Request) (ParserResult, error)  {
	fmt.Println("fetch url", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Println("fetch url error,", err)
		return ParserResult{}, err
	}

	parserResult := r.ParserFunc(body)
	return parserResult, nil
}
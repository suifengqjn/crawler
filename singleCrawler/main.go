package main

import (
	"bufio"
	"fmt"
	"goDemo/Project/crawler/singleCrawler/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	res, err := http.Get(config.CRAWLER_URL)
	if err != err {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK  {
		fmt.Println("StatusCode:%v", res.StatusCode)
		return
	}
	
	
	e := determinEncodeing(res.Body)
	//gbk 转换为 utf8
	utf8Reader := transform.NewReader(res.Body, e.NewDecoder())

	data, err := ioutil.ReadAll(utf8Reader)
	//fmt.Println(string(data))
	printCityList(data)
}

func determinEncodeing (r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e,_, _ := charset.DetermineEncoding(bytes,"")
	return e
}

func printCityList(contents []byte)  {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	for _,m := range matches {
		
		fmt.Printf("City:%s,URL:%s\n", m[2], m[1])
	}
	fmt.Println(len(matches))

}
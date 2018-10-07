package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.zhenai.com/zhenghun")
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
	fmt.Println(string(data))
}

func determinEncodeing (r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e,_, _ := charset.DetermineEncoding(bytes,"")
	return e
}
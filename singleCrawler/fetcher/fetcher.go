package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string)([]byte, error)  {
	res, err := http.Get(url)
	if err != err {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK  {
		fmt.Println("StatusCode:%v", res.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}

	bodyReader := bufio.NewReader(res.Body)
	e := determinEncodeing(bodyReader)
	// 转换为 utf8
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}



func determinEncodeing (r *bufio.Reader) encoding.Encoding  {

	bytes, err := r.Peek(1024)
	if err != nil {
		log.Println("deternime code error")
		return unicode.UTF8
	}
	e,_, _ := charset.DetermineEncoding(bytes,"")
	return e
}

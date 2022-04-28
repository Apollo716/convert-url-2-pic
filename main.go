package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jszwec/csvutil"
)

type Url struct {
	url string `csv:"URL"`
}

const (
	dirPath = "./src"
)

func main() {
	fmt.Println("hogehoge")
	var urls []*Url
	byte, err := ioutil.ReadFile("url.csv")
	log.Println(byte)
	if err != nil {
		log.Printf("readFile err: err is %s", err.Error())
		return
	}
	if err = csvutil.Unmarshal(byte, &urls); err != nil {
		log.Printf("csvUnmarshal err: err is %s", err.Error())
		return
	}
	log.Println(urls[0])
	for i, urlField := range urls {
		var url string = urlField.url
		log.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("fail to get %s: err is %s", url, err.Error())
			return
		}
		defer resp.Body.Close()
		file, err := os.Create(fmt.Sprintf("%s/pic_%d.jpg", dirPath, i))
		if err != nil {
			log.Printf("fileCreate err: err is %s", err.Error())
			return
		}
		defer file.Close()
		io.Copy(file, resp.Body)
	}
}

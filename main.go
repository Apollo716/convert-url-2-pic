package main

import (
	"io/ioutil"
	"log"

	"github.com/jszwec/csvutil"
)

type Url struct {
	url string `csv:"url"`
}

func main() {
	var urls []*Url
	byte, err := ioutil.ReadFile("url.csv")
	if err != nil {
		log.Printf("readFile err: err is %s", err.Error())
		return
	}
	if err = csvutil.Unmarshal(byte, &urls); err != nil {
		log.Printf("csvUnmarshal err: err is %s", err.Error())
		return
	}
}

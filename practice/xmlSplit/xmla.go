package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Member -
type Row struct {
	모집단위   string
	접수인원   int
	선발예정인원 int
	경쟁률    float32
}

//Members -
type root struct {
	Row []Row
}

func main() {
	// xml 파일 오픈
	fp, err := os.Open("nine.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, err := ioutil.ReadAll(fp)

	// xml 디코딩
	var Root root
	xmlerr := xml.Unmarshal(data, &Root)
	if xmlerr != nil {
		panic(xmlerr)
	}

	fmt.Println(Root)
}

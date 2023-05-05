package main

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/ledongthuc/pdf"
	"strings"
)

// todo 指定文件  指定目录 遍历目录 指定关键字  逐行读取 按页读取
//

func main() {
	pdf.DebugOn = true
	content, err := readPdf("demo.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	for word := range KeyWords {
		KeyWords[word] += strings.Count(content, word)
	}

	for key := range KeyWords {
		color.HiGreen("%s 出现次数: %d", key, KeyWords[key])
	}
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

package main

import (
	"strconv"
	"test/myFunc"
	"test/reptile"
)

func main() {
	var imgUrls []string
	var url string
	for i := 1; i < 27; i++ {
		url = `https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/` + strconv.Itoa(i) + ".html"
		imgUrls = append(imgUrls, url)
	}
	reptile.ReptilePic(imgUrls)

	_ = myFunc.PrintSyntaxTrees("D:/GoProjects/test.go")

}

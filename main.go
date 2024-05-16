package main

import (
	"fmt"
	"go/token"
	"test/myFunc"
	"text/scanner"
)

func main() {
	// var imgUrls []string
	// var url string
	// for i := 1; i < 27; i++ {
	// 	url = `https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/` + strconv.Itoa(i) + ".html"
	// 	imgUrls = append(imgUrls, url)
	// }
	// reptile.ReptilePic(imgUrls)11

	_ = myFunc.PrintSyntaxTrees("D:/GoProjects/test.go")

	var src = []byte(`println("你好，世界")`)

	var fset = token.NewFileSet()
	var file = fset.AddFile("D:/GoProjects/test.go", fset.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
	fmt.Println(`ok`)
}

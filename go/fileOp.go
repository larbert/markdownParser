package main

import (
    "os"
    "log"
    "io/ioutil"
    "regexp"
)

const HTMLHeader = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title></title>
</head>
<body>
`

const HTMLFooter = `</body>
</html>`

func main() {
    // 获取文件名
    if len(os.Args) != 2 {
        log.Fatal("参数数量错误")
    }
    fileName := os.Args[1]
    // 打开文件 
    markdownFile, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    // 获取文件内容
    markdownData, err := ioutil.ReadAll(markdownFile)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Data as hex: %x\n", markdownData)
    log.Printf("Data as string: %s\n", markdownData)
    log.Println("Number of bytes read:", len(markdownData))
    // 解析
    regHeader1 := regexp.MustCompile(`# .*`)
    if regHeader1 == nil {
        log.Fatal("regexp err")
    }
    header1Func := func(s string) string{
        return "<h1>" + s[2:] + "</h1>"
    }
    result := regHeader1.ReplaceAllStringFunc(string(markdownData), header1Func)
    log.Println(result)
    // 创建输出文件
    htmlFileName := fileName[:len(fileName)-3] + ".html"
    htmlFile, err := os.Create(htmlFileName)
    if err != nil {
        log.Fatal(err)
    }
    // 写入html
    err = ioutil.WriteFile(htmlFileName, []byte(HTMLHeader + result + HTMLFooter), 0666)
    if err != nil {
        log.Fatal(err)
    }
    // 关闭文件
    defer func(){
        markdownFile.Close()    
        htmlFile.Close()
    }()
}

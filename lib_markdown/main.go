package main

import (
	"bytes"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"html/template"

	"github.com/russross/blackfriday"
)

var h = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }}</title>
</head>
<body>
{{ .Body|unescaped }}
{{ .APIS|unescaped }}
</body>
</html>
`

type Options struct {
	Title string
	Body  string
	APIS  string
}

func unescaped(x string) interface{} { return template.HTML(x) }

func ParseHTML(h string, options *Options) (body string, err error) {
	t := template.New("template.html")
	t.Funcs(template.FuncMap{"unescaped": unescaped})
	t, err = t.Parse(h)

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, options)
	if err != nil {
		return "", err
	}
	body = buf.String()
	return body, nil
}

var table = `
<h3 align="center">项目增补单</h3>
<table border="1" width="800" cellspacing="0" align="center">
	<tr align="center">
    	<td>序号</td>
        <td>维修项目及更换配件</td>
        <td>单价</td>
        <td>数量</td>
        <td>小计</td>
        <td>工时费</td>
        <td>合计</td>
        <td>故障原因</td>
    </tr>
    <tr align="center">
    	<td>1</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr align="center">
    	<td>2</td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
</table>
`

func main() {
	input := []byte("### Markdown\n")

	for i := 1; i <= 5; i++ {
		s := []byte(fmt.Sprintf("#### line %v\n", i))
		input = append(input, s...)
	}

	//input = append(input, []byte(table)...)

	//for i := 1; i <= 5; i++ {
	//	s := []byte(fmt.Sprintf("| %d      | %d       |\n", i, i))
	//	input = append(input, s...)
	//}
	//
	//apisBytes := []byte{}

	output := blackfriday.Run(input)

	html := bluemonday.UGCPolicy().SanitizeBytes(output)

	fmt.Println(string(html))
	options := &Options{
		Title: "Markdown转Html",
		Body:  string(html),
		APIS:  table,
	}
	body, err := ParseHTML(h, options)
	if err != nil {
		panic(err)
	}

	fmt.Println(body)
}

package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		urlFmt string
		err    error
	)
	ctx, cancel = chromedp.NewContext(context.Background())
	defer cancel()
	var buf []byte
	urlFmt = "https://research.swtch.com/gomm"
	//urlFmt = "https://github.com/penk110"
	if err = chromedp.Run(ctx, printToPDF(urlFmt, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("go_memory_model.pdf", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// 生成任务列表
func printToPDF(urlFmt string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlFmt), // 浏览指定的页面
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx) // 通过cdp执行PrintToPDF
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}

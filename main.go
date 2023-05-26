package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func a(msg string) chromedp.ActionFunc {
	return func(context.Context) error {
		fmt.Println(msg)
		return nil
	}
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// fmt.Println("Введите запрос:")
	// scanner := bufio.NewReader(os.Stdin)
	// query, _ := scanner.ReadString('\n')
	var buf0 []byte
	var buf1 []byte
	var buf2 []byte
	var buf3 []byte
	var res string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.youtube.com/"),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.WaitReady("input[name=search_query]", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.Focus("input[name=search_query]"),
		chromedp.Sleep(2*time.Second),
		chromedp.CaptureScreenshot(&buf0),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.WaitReady("input[name=search_query]", chromedp.ByQuery),
		chromedp.SendKeys("input[name=search_query]", "test", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.CaptureScreenshot(&buf1),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.WaitReady("input[name=search_query]", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.KeyEvent(kb.Enter),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Sleep(20*time.Second),
		chromedp.CaptureScreenshot(&buf2),
		chromedp.Text("#video-title", &res, chromedp.ByQuery),
		chromedp.CaptureScreenshot(&buf3),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	save("screen0.png", buf0)
	save("screen1.png", buf1)
	save("screen2.png", buf2)
	save("screen3.png", buf3)
}

func save(pth string, buf []byte) {
	err := ioutil.WriteFile(pth, buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	var query string
	fmt.Println("Введите запрос:")
	fmt.Fscan(os.Stdin, &query)
	var buf1 []byte
	var buf2 []byte
	var buf3 []byte
	var res [3]string
	title := [3]string{
		fmt.Sprintf("#contents > ytd-rich-item-renderer:nth-child(%d) #video-title", 1),
		fmt.Sprintf("#contents > ytd-rich-item-renderer:nth-child(%d) #video-title", 2),
		fmt.Sprintf("#contents > ytd-rich-item-renderer:nth-child(%d) #video-title", 3),
	}
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.youtube.com/"),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.WaitReady("input[name=search_query]", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.Focus("input[name=search_query]"),
		chromedp.SendKeys("input[name=search_query]", query, chromedp.ByQuery),
		chromedp.CaptureScreenshot(&buf1),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.WaitReady("input[name=search_query]", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.KeyEvent(kb.Enter),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Sleep(20*time.Second),
		chromedp.CaptureScreenshot(&buf2),
		chromedp.WaitVisible("ytd-app", chromedp.ByQuery),
		chromedp.Text(title[0], &res[0]),
		chromedp.Text(title[1], &res[1]),
		chromedp.Text(title[2], &res[2]),
		chromedp.CaptureScreenshot(&buf3),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Названия первых трех видео:")
	for i := 0; i < 3; i++ {
		fmt.Println(res[i])
	}
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

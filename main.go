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

func a(num int) chromedp.ActionFunc {
	return func(context.Context) error {
		fmt.Println(num)
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
		chromedp.Sleep(120*time.Second),
		chromedp.Navigate("https://www.youtube.com/"),
		a(1),
		chromedp.Sleep(10*time.Second),
		a(2),
		chromedp.Sleep(60*time.Second),
		chromedp.Focus("input[name=search_query]"),
		chromedp.Sleep(60*time.Second),
		a(3),
		chromedp.SendKeys("input[name=search_query]", query, chromedp.ByQuery),
		a(4),
		chromedp.CaptureScreenshot(&buf1),
		a(5),
		chromedp.Sleep(120*time.Second),
		a(6),
		chromedp.KeyEvent(kb.Enter),
		a(7),
		chromedp.Sleep(120*time.Second),
		a(8),
		chromedp.CaptureScreenshot(&buf2),
		a(9),
		chromedp.Sleep(120*time.Second),
		a(10),
		chromedp.Text(title[0], &res[0]),
		chromedp.Text(title[1], &res[1]),
		chromedp.Text(title[2], &res[2]),
		a(11),
		chromedp.Sleep(120*time.Second),
		a(12),
		chromedp.Screenshot("ytd-app", &buf3, chromedp.ByQuery),
		a(13),
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

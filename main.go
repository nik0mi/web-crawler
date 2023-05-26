package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func foo(context.Context) error {
	fmt.Println("HI")
	return nil
}

func main() {
	var a chromedp.ActionFunc = foo
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// fmt.Println("Введите запрос:")
	// scanner := bufio.NewReader(os.Stdin)
	// query, _ := scanner.ReadString('\n')
	var buf1 []byte
	var buf2 []byte
	var buf3 []byte
	var res string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.youtube.com/"),
		a,
		chromedp.Focus("input[name=search_query]"),
		a,
		chromedp.SendKeys("input[name=search_query]", "test", chromedp.ByQuery),
		a,
		chromedp.CaptureScreenshot(&buf1),
		a,
		chromedp.KeyEvent(kb.Enter),
		a,
		chromedp.WaitVisible("yt-touch-feedback-shape", chromedp.ByQuery),
		a,
		chromedp.CaptureScreenshot(&buf2),
		a,
		chromedp.WaitVisible("ytd-app", chromedp.ByQuery),
		a,
		chromedp.Text("#video-title", &res, chromedp.ByQuery),
		a,
		chromedp.CaptureScreenshot(&buf3),
		a,
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
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

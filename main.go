package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// fmt.Println("Введите запрос:")
	// scanner := bufio.NewReader(os.Stdin)
	// query, _ := scanner.ReadString('\n')
	// fmt.Println(query)
	//var buf []byte
	var res string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.youtube.com/"),
		// chromedp.Text("#video-title", &res, chromedp.ByQuery),
		chromedp.SendKeys("input[name=search_query]", "some"),
		chromedp.KeyEvent(kb.Enter),
		chromedp.Text("#video-title", &res, chromedp.ByQuery),
		// chromedp.Screenshot(".style-scope ytd-masthead", &buf),
		// chromedp.SendKeys("input[name=search_query]", "s"),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	// err := ioutil.WriteFile("screen.png", buf, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

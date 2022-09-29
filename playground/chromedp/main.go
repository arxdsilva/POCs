package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"

// "github.com/chromedp/chromedp"
// )

// func main() {
// 	ctx, ccancel := context.WithTimeout(context.Background(), time.Minute*5)
// 	defer ccancel()

// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// ctx, _ = chromedp.NewContext(ctx)
// fmt.Println("starting: routine ", i)
// go chromedp.Run(context.Background(), chromedp.Navigate("https://twitch.tv/superserverbrasil"), chromedp.ActionFunc(func(ctx context.Context) error {
// 	fmt.Println("go routine ", i)
// 	defer wg.Done()
// 	time.Sleep(time.Minute * 10)
// 	return nil
// }))
// }
// defer cancel()
// time.After(time.Minute)
// 	wg.Wait()
// }

// import (
// 	"fmt"
// 	"log"

// 	"gopkg.in/headzoo/surf.v1"
// )

// func main() {
// 	bow := surf.NewBrowser()
// 	err := bow.Open("http://twitch.tv/superserverbrasil")
// 	if err != nil {
// 		log.Println("error: ", err)
// 	}

// 	// Outputs: "The Go Programming Language"
// 	fmt.Println(bow.Title())
// }

// import (
// 	"fmt"
// 	"sync"
// 	"time"

// 	"github.com/go-rod/rod"
// 	"github.com/go-rod/rod/lib/launcher"
// 	"github.com/go-rod/rod/lib/proto"
// 	// "github.com/go-rod/rod/lib/devices"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	// u := "ws://127.0.0.1:9222/devtools/browser/bd8d2551-ff8d-4f51-8c0f-ddb849366325"
// 	// path, found := launcher.LookPath()
// 	// if !found {
// 	// 	fmt.Println("found: ", found)
// 	// 	return
// 	// }
// 	// u := launcher.New().Bin(path).MustLaunch()
// 	wsURL := launcher.NewUserMode().MustLaunch()
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		fmt.Println("go start: ", i)
// 		go func(id int) {
// 			defer wg.Done()

// 			// page := rod.New().MustConnect().MustPage()
// 			page := rod.New().ControlURL(wsURL).MustConnect().NoDefaultDevice().MustPage()

// 			// err := page.Emulate(devices.Nexus10.Landescape())
// 			// if err != nil {
// 			// 	panic(err)
// 			// }

// 			page.MustNavigate("https://twitch.tv/superserverbrasil")

// 			time.Sleep(time.Second * 10)
// 			page.MustScreenshot(fmt.Sprintf("screenshot%d.png", id))
// 			fmt.Println("go done: ", id)
// 			page.Mouse.Click(proto.InputMouseButton())
// 		}(i)
// 	}

// 	fmt.Println("go waiting")
// 	wg.Wait()
// }

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func init() {
	launcher.NewBrowser().MustGet()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("go start: ", i)
		go func(id int) {
			defer wg.Done()
			browser := rod.New().ControlURL(
				launcher.New().Headless(false).Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome").MustLaunch(),
				// launcher.New().Headless(false).MustLaunch(),
			).MustConnect().MustIncognito()
			defer browser.MustClose()

			page := stealth.MustPage(browser)
			page.MustNavigate("https://twitch.tv/superserverbrasil")
			// page.MustNavigate("https://twitter.com/intz")
			// page.MustNavigate("https://www.youtube.com/watch?v=m01ZeuwC2-M")

			time.Sleep(time.Minute)
			page.MustScreenshot(fmt.Sprintf("screenshot%d.png", id))
		}(i)
	}
	wg.Wait()
}

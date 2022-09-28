package main

import (
	"fmt"
	"log"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/jchv/go-webview2"
)

func main() {
	systray.Register(onReady, nil)

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Webview2",
			Width:  400,
			Height: 600,
			IconId: 2, // icon resource id
			Center: true,
		},
	})
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetSize(400, 600, webview2.HintFixed)
	w.Navigate("http://127.0.0.1:13001")
	w.Run()
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Webview example")
	mShowLantern := systray.AddMenuItem("Show Lantern", "")
	mShowWikipedia := systray.AddMenuItem("Show Wikipedia", "")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mShowLantern.ClickedCh:
				fmt.Println("ClickedCh")
			case <-mShowWikipedia.ClickedCh:
				fmt.Println("ClickedCh")
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

}

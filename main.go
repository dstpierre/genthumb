package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"

	"github.com/chromedp/chromedp"
)

type Data struct {
	Direction       string
	Code            string
	Title           string
	BackgroundColor template.HTML
}

var data Data

func main() {
	flag.StringVar(&data.Direction, "dir", "left", "either left (default) or right")
	flag.StringVar(&data.Code, "code", "", "sample code")
	flag.StringVar(&data.Title, "title", "", "title")
	flag.Parse()

	data.BackgroundColor = template.HTML(pickBackgroundColor())

	go serve()

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, gen(&buf))
	if err != nil {
		fmt.Println("error while taking screenshot: ", err)
		return
	}

	if err = os.WriteFile("thumb.png", buf, 0664); err != nil {
		fmt.Println("error writing output file: ", err)
	}

	os.Exit(0)
}

func gen(buf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("http://localhost:9876/"),
		chromedp.Screenshot("div#thumb", buf),
		//chromedp.FullScreenshot(buf, 100),
	}
}

func serve() {
	http.HandleFunc("/left.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "left.png")
	})
	http.HandleFunc("/right.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "right.png")
	})
	http.HandleFunc("/Anton-Regular.ttf", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Anton-Regular.ttf")
	})
	http.HandleFunc("/dracula.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dracula.css")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("thumb.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println(http.ListenAndServe(":9876", nil))
}

func pickBackgroundColor() string {
	r := rand.Intn(140)
	g := rand.Intn(120)
	b := rand.Intn(140)

	return fmt.Sprintf("%d, %d, %d, 1", r, g, b)
}

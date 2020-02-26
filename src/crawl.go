package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

func main() {

	options := bufio.NewReader(os.Stdin)
	gr := color.New(color.FgHiGreen, color.Bold)
	finflag := make(chan string)
	gr.Printf("Enter URL  : ")
	urlval, _ := options.ReadString('\n')
	urlval = removenewline(urlval)
	go crawlforms(urlval, finflag)
	<-finflag
	go crawllinks(urlval, finflag)
	<-finflag

}

func crawlforms(urlval string, finflag chan string) {
	ylw := color.New(color.FgHiYellow, color.Bold)
	wht := color.New(color.FgHiWhite, color.Bold)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246"),
	)

	c.OnHTML("form", func(e *colly.HTMLElement) {
		action := e.Attr("action")
		id := e.Attr("id")
		method := e.Attr("method")

		ylw.Println("Action : " + action)
		ylw.Println("Form ID : " + id)
		ylw.Println("Method : " + method)
		wht.Println("---------------------------------------------------")
		e.ForEach("input", func(index int, element *colly.HTMLElement) {
			if element.Attr("type") == "text" || element.Attr("type") == "password" {
				ylw.Println("Name ==> " + element.Attr("name") + " [Type : " + element.Attr("type") + "]")
				wht.Println(element.Request.URL.String())
			}
		})
		wht.Println("---------------------------------------------------")

	})
	c.Visit(urlval)
	finflag <- "finished"
}
func crawllinks(urlval string, finflag chan string) {
	bl := color.New(color.FgHiBlue, color.Bold)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		bl.Printf("Link : %q \n", e.Request.AbsoluteURL(e.Attr("href")))

	})
	c.Visit(urlval)
	finflag <- "finished"

}

func removenewline(val string) string {
	if runtime.GOOS == "linux" {
		val = strings.Replace(val, "\n", "", -1)
	} else {
		val = strings.TrimSuffix(val, "\r\n")
	}
	return val
}

func formsubmissions(urlval string) {
	formdata := url.Values{}
	formdata.Set("username", "admin")
	formdata.Set("password", "admin")

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlval, strings.NewReader(formdata.Encode()))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246")
	req.Header.Add("Referer", "https://www.hackthissite.org")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

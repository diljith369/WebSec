package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246"),
	)

	c.OnHTML("form", func(e *colly.HTMLElement) {
		action := e.Attr("action")
		id := e.Attr("id")
		method := e.Attr("method")

		fmt.Println("Action : " + action)
		fmt.Println("Form ID : " + id)
		fmt.Println("Method : " + method)

		e.ForEach("input", func(index int, element *colly.HTMLElement) {
			if element.Attr("type") == "text" || element.Attr("type") == "password" {
				fmt.Println("Name ==> " + element.Attr("name") + " [Type : " + element.Attr("type") + "]")
				fmt.Println(element.Request.URL.String())
			}
		})
	})
	/*err := c.Post("https://www.hackthissite.org/user/login", map[string]string{"username": "admin", "password": "admin"})
	if err != nil {
		fmt.Println(err)
	}
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", string(r.Body))
	})*/

	/*c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Request.AbsoluteURL(e.Attr("href")), e.Attr("href"))

		//fmt.Printf(link)

		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})*/

	// Before making a request print "Visiting ..."
	/*c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})*/

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.hackthissite.org")
	//c2.Visit("https://www.hackthissite.org")

	//link := e.Request.AbsoluteURL(e.Attr("href"))
}

func formsubmissions() {
	formdata := url.Values{}
	formdata.Set("username", "admin")
	formdata.Set("password", "admin")

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://www.hackthissite.org/user/login", strings.NewReader(formdata.Encode()))
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

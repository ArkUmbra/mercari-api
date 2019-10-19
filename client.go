package api

import (
	"fmt"
	// "net/http"
	// "io/ioutil"
	// "errors"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"log"
    "regexp"
)

type Item struct {
	Url string
	Desc string
	Price float64
	Currency string
}

func GetByKeyword(keyword string) (string, error) {
	url := "https://www.mercari.com/jp/search/?page=1&keyword=" + keyword + "&sort_order=&price_max=10000"
	doc, _ := goquery.NewDocument(url)
	// doc.Find("items-box").Each(func(_ int, s *goquery.Selection) {
	// 	 url, _ := s.Attr("href")
	// 	 fmt.Println(url)
	// })
	// var sel = doc.Find("items-box-content").Each(func(_ int, s *goquery.Selection) {
	
	items := []Item{}
	doc.Find(".items-box").Each(func(_ int, s *goquery.Selection) {
		//class, _ := s.Attr("class")
		//fmt.Println(class)
		//fmt.Println("======")

		//href, _ := s.Attr("href")
		//fmt.Println(href)

		href, _ := s.Find("a").Attr("href")
		//fmt.Println(href)

		desc, _ := s.Find(".items-box-name").Html()
		//fmt.Println(desc)

		priceRaw, _ := s.Find(".items-box-price").Html()
		//fmt.Println(priceRaw)

		// Make a Regex to say we only want numbers
		priceRegex, err := regexp.Compile("[^0-9]+")
		if err != nil {
			log.Fatal(err)
		}
	

		var priceStr = priceRegex.ReplaceAllString(priceRaw, "")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Fatal(err)
		}

		var currency = string([]rune(priceRaw)[0:1])

		item := Item{href, desc, price, currency}
		items = append(items, item)
	})

	for _, item := range items {
		fmt.Printf("%+v\n", item)
	}


	return "test", nil

	// var resp, geterr = http.Get(url)

	// if geterr != nil {
	// 	fmt.Println("Error")
	// 	return "", errors.New("Http GET failed")
	// }
	// defer resp.Body.Close()

	// //fmt.Println(resp)
	// //fmt.Println(geterr)

	// var body, readerr = ioutil.ReadAll(resp.Body)
	// if readerr != nil {
	// 	fmt.Println(readerr)
	// 	return "", errors.New("Failed to read response body")
	// }

	// fmt.Println("Printing body", string(body))
	// // fmt.Println(err2)
	// return string(body), nil
}

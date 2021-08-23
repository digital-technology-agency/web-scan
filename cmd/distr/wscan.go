package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/digital-technology-agency/web-scan/pkg/models"
	"github.com/digital-technology-agency/web-scan/pkg/utils"
	"log"
	"net/http"
)

func main() {
	alphabet := flag.String(`alphabet`, "", `Example abcdefg`)
	urlLen := flag.String(`len`, "", `Example 2`)
	flag.Parse()
	/*check flags*/
	if *alphabet == "" && *urlLen == "" {
		flag.PrintDefaults()
		return
	}
	list := map[string]models.Page{}
	for domenName := range models.Gen(*alphabet, utils.Int(*urlLen)) {
		url := fmt.Sprintf("https://%s.ru", domenName)
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Err:[%s]\n", err.Error())
			continue
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
			continue
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			fmt.Printf("Err:[%s]\n", err.Error())
			continue
		}
		item := models.Page{}
		doc.Find("title").Each(func(i int, s *goquery.Selection) {
			item.Title = s.Text()
		})
		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			if s.AttrOr("name", "") == "description" {
				item.Description = s.AttrOr("content", "")
			}
		})
		list[domenName] = item
	}
	println(len(list))
	for key, value := range list {
		fmt.Printf("Domen:[%s]\nTitle:[%s]\nDescription:[%s]", key, value.Title, value.Description)
	}
}

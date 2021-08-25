package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/digital-technology-agency/web-scan/pkg/models"
	"github.com/digital-technology-agency/web-scan/pkg/services"
	"github.com/digital-technology-agency/web-scan/pkg/utils"
	"net/http"
	"sync"
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
	increment := 0
	wg := sync.WaitGroup{}
	for domenName := range services.Gen(*alphabet, utils.Int(*urlLen)) {
		increment += 1
		wg.Add(1)
		go func(domen string, wg *sync.WaitGroup) {
			defer wg.Done()
			url := fmt.Sprintf("https://%s.ru", domen)
			res, err := http.Get(url)
			if err != nil {
				fmt.Printf("Err:[%s]\n", err.Error())
				return
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				fmt.Printf("Status code [%d] error [%s]", res.StatusCode, res.Status)
				return
			}
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				fmt.Printf("Err:[%s]\n", err.Error())
				return
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
			list[domen] = item
		}(domenName, &wg)
	}
	wg.Wait()
	println(fmt.Sprintf("Total size:[%d] Result:[%d]", increment, len(list)))
	for key, value := range list {
		fmt.Printf("Domen:[%s] Title:[%s] Description:[%s]\n", key, value.Title, value.Description)
	}
}

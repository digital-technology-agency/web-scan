package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/digital-technology-agency/web-scan/pkg/models"
	"github.com/digital-technology-agency/web-scan/pkg/services/generators"
	"github.com/digital-technology-agency/web-scan/pkg/services/json"
	"github.com/digital-technology-agency/web-scan/pkg/utils"
	"net/http"
	"os"
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
	total := 0
	domenNames := 0
	wg := sync.WaitGroup{}
	gen := generators.SimpleGenerator{
		Alphabet: *alphabet,
		Len:      utils.Int(*urlLen),
	}
	writer, err := json.NewEachRowWriter("https.txt")
	if err != nil {
		fmt.Printf("Err:[%s]\n", err.Error())
		os.Exit(-1)
	}
	for domenName := range gen.Gen() {
		total += 1
		wg.Add(1)
		go func(domen string, wg *sync.WaitGroup, w *json.EachRowWriter) {
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
			err = w.WriteLine(item)
			if err != nil {
				fmt.Printf("Write line err:[%s]\n", err.Error())
				return
			}
			domenNames += 1
		}(domenName, &wg, writer)
	}
	wg.Wait()
	println(fmt.Sprintf("Total size:[%d] Result:[%d]", total, domenNames))
	/*	for key, value := range list {
		fmt.Printf("Domen:[%s] Title:[%s] Description:[%s]\n", key, value.Title, value.Description)
	}*/
}

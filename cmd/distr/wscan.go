package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/digital-technology-agency/web-scan/pkg/models"
	"github.com/digital-technology-agency/web-scan/pkg/services/generators"
	"github.com/digital-technology-agency/web-scan/pkg/services/json"
	"github.com/digital-technology-agency/web-scan/pkg/utils"
	"github.com/zenthangplus/goccm"
	"net/http"
	"runtime"
)

var (
	coreCount        = flag.String(`core_count`, "1", `Example 1`)
	alphabet         = flag.String(`alphabet`, "", `Example abcdefg`)
	urlLen           = flag.String(`len`, "", `Example 2`)
	concurrencyCount = flag.String(`concurrency`, "10", `Example 10`)
	PROTOCOLS        = []string{"http", "https"}
)

func genWritersProtocols(names []string) map[string]*json.EachRowWriter {
	result := map[string]*json.EachRowWriter{}
	for _, name := range names {
		writer, _ := json.NewEachRowWriter(fmt.Sprintf("%s.txt", name))
		result[name] = writer
	}
	return result
}

func main() {
	flag.Parse()
	/*check flags*/
	if *alphabet == "" && *urlLen == "" {
		flag.PrintDefaults()
		return
	}
	runtime.GOMAXPROCS(utils.Int(*coreCount))
	cuncurency := goccm.New(utils.Int(*concurrencyCount))
	total := 0
	domenNames := 0
	gen := generators.SimpleGenerator{
		Alphabet: *alphabet,
		Len:      utils.Int(*urlLen),
	}
	protocolWriters := genWritersProtocols(PROTOCOLS)
	for domenName := range gen.Gen() {
		cuncurency.Wait()
		total += 1
		for _, protokol := range PROTOCOLS {
			go func(protokol, domen string, w *json.EachRowWriter) {
				defer cuncurency.Done()
				url := fmt.Sprintf("%s://%s.ru", protokol, domen)
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
			}(protokol, domenName, protocolWriters[protokol])
		}
	}
	cuncurency.WaitAllDone()
	println(fmt.Sprintf("Total size:[%d] Result:[%d]", total, domenNames))
	/*	for key, value := range list {
		fmt.Printf("Domen:[%s] Title:[%s] Description:[%s]\n", key, value.Title, value.Description)
	}*/
}

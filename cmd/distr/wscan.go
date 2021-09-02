package main

import (
	"flag"
	"fmt"
	"github.com/digital-technology-agency/web-scan/pkg/services/generators"
	"github.com/digital-technology-agency/web-scan/pkg/services/json"
	"github.com/digital-technology-agency/web-scan/pkg/services/page"
	"github.com/digital-technology-agency/web-scan/pkg/utils"
	"github.com/zenthangplus/goccm"
	"runtime"
)

var (
	coreCount        = flag.String(`core_count`, "1", `Example 1`)
	alphabet         = flag.String(`alphabet`, "", `Example abcdefg`)
	urlLen           = flag.String(`len`, "", `Example 2`)
	concurrencyCount = flag.String(`concurrency`, "5", `Example 5`)
	protocols        = []string{"http", "https"}
)

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
	protocolWriters := json.NewEachRowWriters(protocols)
	for domenName := range gen.Gen() {
		cuncurency.Wait()
		total += 1
		for _, protokol := range protocols {
			go func(protokol, domen string, w *json.EachRowWriter) {
				defer cuncurency.Done()
				url := fmt.Sprintf("%s://%s.ru", protokol, domen)
				pageService := page.PageService{
					Url: url,
				}
				item, err := pageService.ReadPage()
				if err != nil {
					return
				}
				if item == nil {
					fmt.Printf("Page is nil\n")
					return
				}
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
}

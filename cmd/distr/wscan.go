package main

import (
	"flag"
	"fmt"
	"github.com/digital-technology-agency/web-scan/pkg/database"
	"github.com/digital-technology-agency/web-scan/pkg/database/sqlite"
	"github.com/digital-technology-agency/web-scan/pkg/models"
	"github.com/digital-technology-agency/web-scan/pkg/services/generators"
	"github.com/digital-technology-agency/web-scan/pkg/services/page"
	"github.com/digital-technology-agency/web-scan/pkg/utils"
	"github.com/zenthangplus/goccm"
	"os"
	"runtime"
)

var (
	coreCount        = flag.String(`core_count`, "1", `Example 1`)
	alphabet         = flag.String(`alphabet`, "", `Example abcdefg`)
	urlLen           = flag.String(`len`, "", `Example 2`)
	concurrencyCount = flag.String(`concurrency`, "5", `Example 5`)
	protocols        = []string{"http", "https"}
	sqliteService    = sqlite.SqLite{}
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
	/*protocolWriters := json.NewEachRowWriters(protocols)*/
	model := models.Page{}
	err := model.CreateTable(sqliteService)
	if err != nil {
		fmt.Printf("Db service, create table! err:[%s]\n", err.Error())
		os.Exit(-1)
	}
	for domenName := range gen.Gen() {
		cuncurency.Wait()
		total += 1
		for _, protokol := range protocols {
			go func(protokol, domen string, dbService database.DbService) {
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
				err = item.AddOrUpdate(dbService)
				if err != nil {
					fmt.Printf("Write line to service! Err:[%s]\n", err.Error())
					return
				}
				/*				err = w.WriteLine(item)*/
				domenNames += 1
			}(protokol, domenName, sqliteService)
		}
	}
	cuncurency.WaitAllDone()
	println(fmt.Sprintf("Total size:[%d] Result:[%d]", total, domenNames))
}

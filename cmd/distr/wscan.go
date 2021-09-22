package main

import (
	"flag"
	"fmt"
	"github.com/digital-technology-agency/web-scan/pkg/env"
	"github.com/digital-technology-agency/web-scan/pkg/models"
	"github.com/digital-technology-agency/web-scan/pkg/services/generators"
	"github.com/digital-technology-agency/web-scan/pkg/services/json"
	"github.com/digital-technology-agency/web-scan/pkg/services/page"
	"github.com/digital-technology-agency/web-scan/pkg/utils"
	"github.com/zenthangplus/goccm"
	"os"
	"runtime"
)

var (
	processCount     = flag.String(`process_count`, "1", `Example 1`)
	alphabet         = flag.String(`alphabet`, "", `Example abcdefg`)
	urlLen           = flag.String(`len`, "", `Example 2`)
	concurrencyCount = flag.String(`concurrency`, "5", `Example 5`)
	dataStore        = flag.String(`data_store`, env.SQLITE_STORE, "Example:\n"+
		fmt.Sprintf("%s\n", env.SQLITE_STORE)+
		fmt.Sprintf("%s\n", env.JSON_EACH_ROW_STORE)+
		"",
	)
	protocols = []string{env.HTTP_PROTOCOL, env.HTTPS_PROTOCOL}
	/*services*/
	protocolWriters = map[string]*json.EachRowWriter{}
	dbStore         = env.InitDbStore()
)

func main() {
	flag.Parse()
	/*check flags*/
	if *alphabet == "" && *urlLen == "" {
		flag.PrintDefaults()
		return
	}
	if !env.CheckStore(*dataStore) {
		fmt.Printf("Store [%s] - not found!\n", *dataStore)
		flag.PrintDefaults()
		return
	}
	runtime.GOMAXPROCS(utils.Int(*processCount))
	cuncurency := goccm.New(utils.Int(*concurrencyCount))
	total := 0
	domenNames := 0
	gen := generators.SimpleGenerator{
		Alphabet: *alphabet,
		Len:      utils.Int(*urlLen),
	}
	model := models.Page{}
	/*check db service*/
	if env.CheckDataBaseStore(*dataStore) {
		err := model.CreateTable(dbStore[*dataStore])
		if err != nil {
			fmt.Printf("Db service, create table! err:[%s]\n", err.Error())
			os.Exit(-1)
		}
	} else {
		protocolWriters = json.NewEachRowWriters(protocols)
	}
	for domenName := range gen.Gen() {
		cuncurency.Wait()
		total += 1
		for _, protokol := range protocols {
			go func(protokol, domen string) {
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
				switch *dataStore {
				default:
					fmt.Printf("Store [%s] - not found!\n", *dataStore)
					break
				case env.JSON_EACH_ROW_STORE:
					err = protocolWriters[protokol].WriteLine(item)
					break
				case env.SQLITE_STORE:
					err = item.AddOrUpdate(dbStore[env.SQLITE_STORE])
					break
				}
				if err != nil {
					fmt.Printf("Write line to service! Err:[%s]\n", err.Error())
					return
				}
				domenNames += 1
			}(protokol, domenName)
		}
	}
	cuncurency.WaitAllDone()
	println(fmt.Sprintf("Total size:[%d] Result:[%d]", total, domenNames))
}

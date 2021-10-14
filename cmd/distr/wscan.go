package main

import (
	"flag"
	"fmt"
	"github.com/digital-technology-agency/web-scan/pkg/config"
	"github.com/digital-technology-agency/web-scan/pkg/database"
	"github.com/digital-technology-agency/web-scan/pkg/models"
	"github.com/digital-technology-agency/web-scan/pkg/services/page"
	"github.com/zenthangplus/goccm"
	"os"
	"runtime"
)

var (
	configurationFile = flag.String(`configuration_file`, "", `Example: config.json`)
	configuration     = config.Default()
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "init" {
			err := configuration.Save("config.json")
			if err != nil {
				fmt.Printf("Configuration file not save! err:[%s]\n", err.Error())
				os.Exit(-1)
			}
			os.Exit(0)
		}
	}
	flag.Parse()
	if *configurationFile != "" {
		loadConfig, err := config.Load(*configurationFile)
		if err != nil {
			fmt.Printf("Configuration file not correct! err:[%s]\n", err.Error())
			os.Exit(-1)
		}
		/*validate config*/
		err = loadConfig.Validate()
		if err != nil {
			fmt.Printf("Configuration file not correct! err:[%s]\n", err.Error())
			os.Exit(-1)
		}
		configuration = *loadConfig
	}
	configuration.InitGenerator()
	configuration.InitDataStore()
	runtime.GOMAXPROCS(configuration.ProcessCount)
	cuncurency := goccm.New(configuration.ConcurrencyCount)
	total := 0
	domenNames := 0
	gen := configuration.Generator
	protocols := configuration.ProtocolTypes
	model := models.Page{}
	/*check db service*/
	err := model.CreateTable(configuration.DataStore)
	if err != nil {
		fmt.Printf("Db service, create table! err:[%s]\n", err.Error())
		os.Exit(-1)
	}
	for domenName := range gen.Gen() {
		cuncurency.Wait()
		total += 1
		for _, protokol := range protocols {
			go func(protokol, domen string, dataStore database.DbService) {
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
				err = item.AddOrUpdate(dataStore)
				if err != nil {
					fmt.Printf("Write line to service! Err:[%s]\n", err.Error())
					return
				}
				domenNames += 1
			}(protokol, domenName, configuration.DataStore)
		}
	}
	cuncurency.WaitAllDone()
	println(fmt.Sprintf("Total size:[%d] Result:[%d]", total, domenNames))
}

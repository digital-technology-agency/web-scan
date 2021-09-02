package page

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/digital-technology-agency/web-scan/pkg/models"
	"io/ioutil"
	"net/http"
)

// PageService page service.
type PageService struct {
	Url string
}

// ReadPage read page.
func (s PageService) ReadPage() (*models.Page, error) {
	url := s.Url
	item := models.Page{
		Url: url,
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Err:[%s]\n", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("Status code [%d] error [%s]", res.StatusCode, res.Status)
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Printf("Err:[%s]\n", err.Error())
		return nil, err
	}
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		item.Title = s.Text()
	})
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("name", "") == "description" {
			item.Description = s.AttrOr("content", "")
		}
	})
	urlRobotTxt := fmt.Sprintf("%s/robots.txt", url)
	resRobots, err := http.Get(urlRobotTxt)
	if err != nil {
		fmt.Printf("Err:[%s]\n", err.Error())
		return nil, err
	}
	if resRobots.StatusCode != 200 {
		fmt.Printf("Robots txt. Status:[%d]\n", resRobots.StatusCode)
		return nil, err
	}
	allBytesRobotsTxt, err := ioutil.ReadAll(resRobots.Body)
	if err != nil {
		fmt.Printf("Err:[%s]\n", err.Error())
		return nil, err
	}
	item.Robots = string(allBytesRobotsTxt)
	return &item, nil
}

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
	robotsTxtData, err := getRobotsTxt(url)
	if err != nil {
		fmt.Printf("Robots txt err:[%s]\n", err.Error())
		return nil, err
	}
	item.Robots = string(robotsTxtData)
	siteMapXml, err := getSitemapXml(url)
	if err != nil {
		fmt.Printf("Sitemap xml err:[%s]\n", err.Error())
		return nil, err
	}
	item.Sitemap = string(siteMapXml)
	return &item, nil
}

func getSitemapXml(url string) ([]byte, error) {
	urlSitemapXml := fmt.Sprintf("%s/sitemap.xml", url)
	return getResponse(urlSitemapXml)
}

func getRobotsTxt(url string) ([]byte, error) {
	urlRobotTxt := fmt.Sprintf("%s/robots.txt", url)
	return getResponse(urlRobotTxt)
}

func getResponse(url string) ([]byte, error) {
	resRobots, err := http.Get(url)
	if err != nil {
		fmt.Printf("Err:[%s]\n", err.Error())
		return nil, err
	}
	if resRobots.StatusCode != 200 {
		fmt.Printf("Get response data. Status:[%d]\n", resRobots.StatusCode)
		return nil, err
	}
	return ioutil.ReadAll(resRobots.Body)
}

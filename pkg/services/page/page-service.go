package page

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"github.com/digital-technology-agency/web-scan/pkg/models"
)

// Page page service.
type Page struct {
	URL string
}

// ReadPage read page.
func (s Page) ReadPage() (*models.Page, error) {
	url := s.URL
	item := models.Page{
		URL: url,
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
	siteMapXML, err := getSitemapXML(url)
	if err != nil {
		fmt.Printf("Sitemap xml err:[%s]\n", err.Error())
		return nil, err
	}
	item.Sitemap = string(siteMapXML)
	return &item, nil
}

func getSitemapXML(url string) ([]byte, error) {
	urlSitemapXML := fmt.Sprintf("%s/sitemap.xml", url)
	return getResponse(urlSitemapXML)
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

package crawler

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/kanzitelli/good-news-backend/models"
	"github.com/kanzitelli/good-news-backend/utils"
)

// TheVillage <struct>
// is used to present The Village crawler.
type TheVillage struct{}

const (
	baseURL_TV           = "https://www.the-village.ru"
	crawlURL_News_TP     = "https://www.the-village.ru/news"
	crawlURL_Business_TP = "https://www.the-village.ru/village/business"
	crawlURL_Style_TP    = "https://www.the-village.ru/village/service-shopping"
)

// Run <function>
// is used to start crawling process.
func (tv TheVillage) Run() []models.News {
	fmt.Println("Crawling TheVillage...")

	var totalNews []models.News
	newsFuncs := []NewsFunc{
		tv.runNews,
		tv.runBusiness,
		tv.runStyle,
	}

	for _, f := range newsFuncs {
		tmpNews := f()
		totalNews = append(totalNews, tmpNews...)
	}

	return totalNews
}

func (tv TheVillage) runNews() []models.News {
	// creating simple colly instance without any options
	newsCollector := colly.NewCollector()

	// array of news that will be returned
	var news []models.News

	newsCollector.OnHTML(".content", func(divContent *colly.HTMLElement) {
		divContent.ForEach(".post-item-microformat", func(i1 int, divItem *colly.HTMLElement) {
			link := divItem.ChildAttr("a[href]", "href")
			fullLink := fmt.Sprintf("%s%s", baseURL_TV, link)
			title := divItem.ChildText(".post-title")

			_id := utils.MakeHash(fullLink) // here we are going to create hash from full link in order to set ID of a news to hash value, so mongo won't add in case of duplicates

			news = append(news, models.News{
				ID:         _id,
				Title:      title,
				Link:       fullLink,
				Preamble:   "",
				TimeAdded:  time.Now().Unix(),
				NewsType:   models.TypePost,
				NewsSource: models.TheVillageNewsSource,
			})
		})

		divContent.ForEach(".post-item-news", func(i1 int, divItem *colly.HTMLElement) {
			link := divItem.ChildAttr("a[href]", "href")
			fullLink := fmt.Sprintf("%s%s", baseURL_TV, link)
			title := divItem.ChildText(".post-title")

			_id := utils.MakeHash(fullLink) // here we are going to create hash from full link in order to set ID of a news to hash value, so mongo won't add in case of duplicates

			news = append(news, models.News{
				ID:         _id,
				Title:      title,
				Link:       fullLink,
				Preamble:   "",
				TimeAdded:  time.Now().Unix(),
				NewsType:   models.TypeNews,
				NewsSource: models.TheVillageNewsSource,
			})
		})
	})

	newsCollector.Visit(crawlURL_News_TP)
	newsCollector.Wait()

	return news
}

func (tv TheVillage) runBusiness() []models.News {
	// creating simple colly instance without any options
	newsCollector := colly.NewCollector()

	// array of news that will be returned
	var news []models.News

	newsCollector.OnHTML(".row-cont", func(divContent *colly.HTMLElement) {
		divContent.ForEach(".post-block-featured", func(i1 int, divItem *colly.HTMLElement) {
			link := divItem.ChildAttr("a[href]", "href")
			fullLink := fmt.Sprintf("%s%s", baseURL_TV, link)
			title := divItem.ChildText(".post-title")
			preamble := divItem.ChildText(".post-preamble")

			_id := utils.MakeHash(fullLink) // here we are going to create hash from full link in order to set ID of a news to hash value, so mongo won't add in case of duplicates

			news = append(news, models.News{
				ID:         _id,
				Title:      title,
				Link:       fullLink,
				Preamble:   preamble,
				TimeAdded:  time.Now().Unix(),
				NewsType:   models.TypeBusiness,
				NewsSource: models.TheVillageNewsSource,
			})
		})
	})

	newsCollector.Visit(crawlURL_Business_TP)
	newsCollector.Wait()

	return news
}

func (tv TheVillage) runStyle() []models.News {
	// creating simple colly instance without any options
	newsCollector := colly.NewCollector()

	// array of news that will be returned
	var news []models.News

	newsCollector.OnHTML(".row-cont", func(divContent *colly.HTMLElement) {
		divContent.ForEach(".post-block-featured", func(i1 int, divItem *colly.HTMLElement) {
			link := divItem.ChildAttr("a[href]", "href")
			fullLink := fmt.Sprintf("%s%s", baseURL_TV, link)
			title := divItem.ChildText(".post-title")
			preamble := divItem.ChildText(".post-preamble")

			_id := utils.MakeHash(fullLink) // here we are going to create hash from full link in order to set ID of a news to hash value, so mongo won't add in case of duplicates

			news = append(news, models.News{
				ID:         _id,
				Title:      title,
				Link:       fullLink,
				Preamble:   preamble,
				TimeAdded:  time.Now().Unix(),
				NewsType:   models.TypeStyle,
				NewsSource: models.TheVillageNewsSource,
			})
		})
	})

	newsCollector.Visit(crawlURL_Style_TP)
	newsCollector.Wait()

	return news
}

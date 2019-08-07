package crawler

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/kanzitelli/good-news-backend/models"
	"github.com/kanzitelli/good-news-backend/utils"
)

// SecretMag <struct>
// is used to present Secret Magazine crawler.
type SecretMag struct{}

const (
	baseURL_SM  = "https://secretmag.ru"
	crawlURL_SM = "https://secretmag.ru/news"
)

// Run <function>
// is used to start crawling process.
func (sm SecretMag) Run() []models.News {
	fmt.Println("Crawling SecretMag...")

	var totalNews []models.News
	newsFuncs := []NewsFunc{
		sm.runNews,
	}

	for _, f := range newsFuncs {
		tmpNews := f()
		totalNews = append(totalNews, tmpNews...)
	}

	return totalNews
}

func (sm SecretMag) runNews() []models.News {
	// creating simple colly instance without any options
	newsCollector := colly.NewCollector()

	// array of news that will be returned
	var news []models.News

	newsCollector.OnHTML(".wrapper", func(divWrapper *colly.HTMLElement) {
		divWrapper.ForEach(".container", func(i1 int, divContainer *colly.HTMLElement) {
			divContainer.ForEach(".item", func(i2 int, divItem *colly.HTMLElement) {
				link := divItem.ChildAttr("a[href]", "href")
				fullLink := fmt.Sprintf("%s%s", baseURL_SM, link)
				title := divItem.ChildText(".headline")

				_id := utils.MakeHash(fullLink) // here we are going to create hash from full link in order to set ID of a news to hash value, so mongo won't add in case of duplicates

				news = append(news, models.News{
					ID:         _id,
					Title:      title,
					Link:       fullLink,
					Preamble:   "",
					TimeAdded:  time.Now().Unix(),
					NewsType:   models.TypeNews,
					NewsSource: models.SecretMagNewsSource,
				})
			})
		})
	})

	newsCollector.Visit(crawlURL_SM)
	newsCollector.Wait()

	return news
}

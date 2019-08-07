package crawler

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/kanzitelli/good-news-backend/models"
	"github.com/kanzitelli/good-news-backend/utils"
)

// TheoryAndPractice <struct>
// is used to present Secret Magazine crawler.
type TheoryAndPractice struct{}

const (
	baseURL_TP  = "https://theoryandpractice.ru"
	crawlURL_TP = "https://theoryandpractice.ru/posts"
)

// Run <function>
// is used to start crawling process.
func (tp TheoryAndPractice) Run() []models.News {
	fmt.Println("Crawling TheoryAndPractice...")

	var totalNews []models.News
	newsFuncs := []NewsFunc{
		tp.runPosts,
	}

	for _, f := range newsFuncs {
		tmpNews := f()
		totalNews = append(totalNews, tmpNews...)
	}

	return totalNews
}

func (tp TheoryAndPractice) runPosts() []models.News {
	// creating simple colly instance without any options
	newsCollector := colly.NewCollector()

	// array of news that will be returned
	var news []models.News

	newsCollector.OnHTML("div[role]", func(divPostsContainer *colly.HTMLElement) {
		// checking if value of div attribure `role` equals to "postsContainer"
		if divPostsContainer.Attr("role") == "postsContainer" {
			divPostsContainer.ForEach(".preview-box-post", func(i1 int, divPost *colly.HTMLElement) {
				link := divPost.ChildAttr("a[href]", "href")
				fullLink := fmt.Sprintf("%s%s", baseURL_TP, link)
				title := divPost.ChildText(".preview-box-post-title")
				preamble := divPost.ChildText(".preview-box-post-descr")

				_id := utils.MakeHash(fullLink) // here we are going to create hash from full link in order to set ID of a news to hash value, so mongo won't add in case of duplicates

				news = append(news, models.News{
					ID:         _id,
					Title:      title,
					Link:       fullLink,
					Preamble:   preamble,
					TimeAdded:  time.Now().Unix(),
					NewsType:   models.TypePost,
					NewsSource: models.TheoryAndPracticeNewsSource,
				})
			})
		}
	})

	newsCollector.Visit(crawlURL_TP)
	newsCollector.Wait()

	return news
}

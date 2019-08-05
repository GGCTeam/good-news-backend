package crawler

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/kanzitelli/good-news-backend/models"
)

// TheoryAndPractice <struct>
// is used to present Secret Magazine crawler.
type TheoryAndPractice struct{}

const (
	baseURLTP  = "https://theoryandpractice.ru"
	crawlURLTP = "https://theoryandpractice.ru/posts"
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
				title := divPost.ChildText(".preview-box-post-title")
				preamble := divPost.ChildText(".preview-box-post-descr")

				news = append(news, models.News{
					Title:      title,
					Link:       fmt.Sprintf("%s%s", baseURLTP, link),
					Preamble:   preamble,
					TimeAdded:  time.Now().Unix(),
					NewsType:   models.TypePost,
					NewsSource: models.TheoryAndPracticeNewSource,
				})
			})
		}
	})

	newsCollector.Visit(crawlURLTP)
	newsCollector.Wait()

	fmt.Println(news)

	return news
}

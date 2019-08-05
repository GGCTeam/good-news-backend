package crawler

import (
	"time"

	"github.com/kanzitelli/good-news-backend/models"
)

// NewsCrawler <interface>
// is used to describe news crawler class instance
type NewsCrawler interface {
	Run() []models.News
}

// NewsFunc <type>
// is used to simplify news func type signature
type NewsFunc func() []models.News

// Start <function>
// is used to start process of web resources crawling every 3 minutes
func Start() {
	go startCrawler()
}

func startCrawler() {
	// array of crawlers for different news sources which implement NewsCrawler interface.
	crawlers := []NewsCrawler{
		SecretMag{},
		TheoryAndPractice{},
		TheVillage{},
	}

	// duration of each crawling process
	duration := time.Second * 5 // time.Minute * 3

	for range time.Tick(duration) {
		// all news collected from each crawler
		var totalNews []models.News

		for _, cr := range crawlers {
			totalNews = append(totalNews, cr.Run()...)
			// fmt.Println(totalNews)

			// here we run all crawlers
			// get all new news
			// then trying to insert them to mongo db
		}
	}
}

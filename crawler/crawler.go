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

// Start <function>
// is used to start process of web resources crawling every 3 minutes
func Start() {
	go startCrawler()
}

func startCrawler() {
	crawlers := []NewsCrawler{
		SecretMag{},
		TheoryAndPractice{},
	}
	duration := time.Second * 5 // time.Minute * 3

	for range time.Tick(duration) {
		// fmt.Println("go and do crawling")
		for _, cr := range crawlers {
			cr.Run()
		}
	}
}

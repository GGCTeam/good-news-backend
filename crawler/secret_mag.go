package crawler

import (
	"fmt"
	"time"

	"github.com/kanzitelli/good-news-backend/models"
)

// SecretMag <struct>
// is used to present Secret Magazine crawler.
type SecretMag struct{}

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
	fmt.Println("Crawling SecretMag News...")
	time.Sleep(time.Millisecond * 200)

	return []models.News{}
}

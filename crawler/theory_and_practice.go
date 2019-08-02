package crawler

import (
	"fmt"
	"time"

	"github.com/kanzitelli/good-news-backend/models"
)

// TheoryAndPractice <struct>
// is used to present Secret Magazine crawler.
type TheoryAndPractice struct{}

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
	fmt.Println("Crawling TheoryAndPractice Posts...")
	time.Sleep(time.Millisecond * 200)

	return []models.News{}
}

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
	time.Sleep(time.Second * 1)

	return []models.News{}
}

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
	time.Sleep(time.Second * 1)

	return []models.News{}
}

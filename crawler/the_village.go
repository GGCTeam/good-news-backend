package crawler

import (
	"fmt"
	"time"

	"github.com/kanzitelli/good-news-backend/models"
)

// TheVillage <struct>
// is used to present The Village crawler.
type TheVillage struct{}

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
	fmt.Println("Crawling TheVillage News...")
	time.Sleep(time.Millisecond * 200)

	return []models.News{}
}

func (tv TheVillage) runBusiness() []models.News {
	fmt.Println("Crawling TheVillage Business...")
	time.Sleep(time.Millisecond * 200)

	return []models.News{}
}

func (tv TheVillage) runStyle() []models.News {
	fmt.Println("Crawling TheVillage Style...")
	time.Sleep(time.Millisecond * 200)

	return []models.News{}
}

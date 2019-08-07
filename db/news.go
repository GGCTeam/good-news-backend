package db

import (
	"github.com/kanzitelli/good-news-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	newsCollection = "news"
)

// NewsInsert <function>
// is used to insert an array of news to DB
func (s Service) NewsInsert(news []models.News) error {
	c := s.Collection(newsCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	// we cast it to array of interfaces because mongo driver asks so
	var iNews []interface{} // []models.NewsType
	iNews = append(iNews, news)

	_, err := c.InsertMany(
		ctx,
		iNews,
		s.InsertManyOptionsOrdered(),
	)

	return err
}

// NewsGet <function>
// is used to get news with given count
func (s Service) NewsGet(count int) ([]models.News, error) {
	c := s.Collection(newsCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	var news []models.News
	cur, err := c.Find(ctx, bson.D{{}})
	if err != nil {
		return []models.News{}, err
	}

	for cur.Next(ctx) {
		var n models.News
		err = cur.Decode(&n)
		if err != nil {
			return []models.News{}, err
		}

		news = append(news, n)
	}

	if err = cur.Err(); err != nil {
		return []models.News{}, err
	}

	cur.Close(ctx)

	return news, nil
}

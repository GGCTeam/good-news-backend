package db

import (
	"github.com/kanzitelli/good-news-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	iNews := []interface{}{} // []models.News
	for _, n := range news {
		iNews = append(iNews, n)
	}

	_, err := c.InsertMany(
		ctx,
		iNews,
		s.InsertManyOptionsOrdered(),
	)

	return err
}

// NewsGet <function>
// is used to get news with given count
func (s Service) NewsGet(count int64) ([]models.News, error) {
	c := s.Collection(newsCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	var news []models.News
	options := options.Find()
	cur, err := c.Find(
		ctx,
		bson.D{},
		options.SetSort(bson.D{{"time_added", -1}}), // desc sort by time_added
		options.SetLimit(count),                     // limit to <count>
	)
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

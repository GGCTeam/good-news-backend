package db

import (
	"github.com/kanzitelli/good-news-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	newsSourcesCollection = "news_sources"
)

// FillNewsSources <function>
// is used to fill news sources in DB
func (s Service) fillNewsSources() error {
	c := s.Collection(newsSourcesCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	// we cast it to array of interfaces because mongo driver asks so
	var newsSources []interface{} // []models.NewsSource
	newsSources = append(
		newsSources,
		models.SecretMagNewsSource,
		models.TheVillageNewsSource,
		models.TheoryAndPracticeNewsSource,
	)

	_, err := c.InsertMany(
		ctx,
		newsSources,
		s.InsertManyOptionsOrdered(),
	)

	return err
}

// NewsSourcesGet <function>
// is used to get news sources
func (s Service) NewsSourcesGet() ([]models.NewsSource, error) {
	c := s.Collection(newsSourcesCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	var newsSources []models.NewsSource
	cur, err := c.Find(ctx, bson.D{})
	if err != nil {
		return []models.NewsSource{}, err
	}

	for cur.Next(ctx) {
		var n models.NewsSource
		err = cur.Decode(&n)
		if err != nil {
			return []models.NewsSource{}, err
		}

		newsSources = append(newsSources, n)
	}

	if err = cur.Err(); err != nil {
		return []models.NewsSource{}, err
	}

	cur.Close(ctx)

	return newsSources, nil
}

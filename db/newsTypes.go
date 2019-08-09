package db

import (
	"github.com/kanzitelli/good-news-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	newsTypesCollection = "news_types"
)

// FillNewsTypes <function>
// is used to fill news types in DB
func (s Service) fillNewsTypes() error {
	c := s.Collection(newsTypesCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	// we cast it to array of interfaces because mongo driver asks so
	var newsTypes []interface{} // []models.NewsType
	newsTypes = append(
		newsTypes,
		models.TypeNews,
		models.TypeBusiness,
		models.TypePost,
		models.TypeStyle,
	)

	_, err := c.InsertMany(
		ctx,
		newsTypes,
		s.InsertManyOptionsOrdered(),
	)

	return err
}

// NewsTypesGet <function>
// is used to get news types
func (s Service) NewsTypesGet() ([]models.NewsType, error) {
	c := s.Collection(newsTypesCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	var newsTypes []models.NewsType
	cur, err := c.Find(ctx, bson.D{})
	if err != nil {
		return []models.NewsType{}, err
	}

	for cur.Next(ctx) {
		var n models.NewsType
		err = cur.Decode(&n)
		if err != nil {
			return []models.NewsType{}, err
		}

		newsTypes = append(newsTypes, n)
	}

	if err = cur.Err(); err != nil {
		return []models.NewsType{}, err
	}

	cur.Close(ctx)

	return newsTypes, nil
}

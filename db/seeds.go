package db

import "github.com/kanzitelli/good-news-backend/models"

const (
	newsSourcesCollection = "news_sources"
	newsTypesCollection   = "news_types"
)

// FillSeedsInformation <function>
// is used to fill information that needed before usage of news endpoint
func (s Service) FillSeedsInformation() {
	s.fillNewsSources()
	s.fillNewsTypes()
}

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

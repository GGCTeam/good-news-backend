package models

var (
	SecretMagNewSource = NewsSource{
		Name:     "Secret Mag",
		ImageURL: "/images/secret_mag.png",
		Types:    []NewsType{TypeNews},
	}

	TheVillageNewSource = NewsSource{
		Name:     "The Village",
		ImageURL: "/images/the_village.png",
		Types:    []NewsType{TypeNews, TypeBusiness, TypeStyle},
	}

	TheoryAndPracticeNewSource = NewsSource{
		Name:     "Theory and Practice",
		ImageURL: "/images/theory_and_practice.png",
		Types:    []NewsType{TypePost},
	}
)

// NewsSource <model>
// is used to describe news source model.
type NewsSource struct {
	Name     string     `json:"name" bson:"name"`
	ImageURL string     `json:"image_url" bson:"image_url"`
	Types    []NewsType `json:"types" bson:"types"`
}

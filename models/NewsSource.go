package models

var (
	SecretMagNewsSource = NewsSource{
		ID:       "f3fefa39-8adf-4e66-8f7c-18a071772a47",
		Name:     "Secret Mag",
		ImageURL: "/images/secret_mag.png",
		Types:    []NewsType{TypeNews},
	}

	TheVillageNewsSource = NewsSource{
		ID:       "5c118e2b-1ed4-41af-bcb3-f5bbc11d55d3",
		Name:     "The Village",
		ImageURL: "/images/the_village.png",
		Types:    []NewsType{TypeNews, TypeBusiness, TypeStyle},
	}

	TheoryAndPracticeNewsSource = NewsSource{
		ID:       "04e00787-598a-40bf-baf6-aa2ba1f8203c",
		Name:     "Theory and Practice",
		ImageURL: "/images/theory_and_practice.png",
		Types:    []NewsType{TypePost},
	}
)

// NewsSource <model>
// is used to describe news source model.
type NewsSource struct {
	ID       string     `json:"_id" bson:"_id"`
	Name     string     `json:"name" bson:"name"`
	ImageURL string     `json:"image_url" bson:"image_url"`
	Types    []NewsType `json:"types" bson:"types"`
}

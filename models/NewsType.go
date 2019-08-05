package models

var (
	TypeNews = NewsType{
		Type: "news",
		Name: "Новости",
	}

	TypePost = NewsType{
		Type: "post",
		Name: "Пост",
	}

	TypeBusiness = NewsType{
		Type: "business",
		Name: "Бизнес",
	}

	TypeStyle = NewsType{
		Type: "style",
		Name: "Стиль",
	}
)

// NewsType <struct>
// is used to describe news type model.
type NewsType struct {
	Type string `json:"type" bson:"type"`
	Name string `json:"name" bson:"name"`
}

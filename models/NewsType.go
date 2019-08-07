package models

var (
	TypeNews = NewsType{
		ID:   "6cef0d43-c49a-401a-9820-eda43a483a21",
		Type: "news",
		Name: "Новости",
	}

	TypePost = NewsType{
		ID:   "f98332cc-37a9-48c8-9abf-12b84de3d673",
		Type: "post",
		Name: "Пост",
	}

	TypeBusiness = NewsType{
		ID:   "ce0ef1ab-655b-475e-98b3-7c8fec74a3e8",
		Type: "business",
		Name: "Бизнес",
	}

	TypeStyle = NewsType{
		ID:   "0c05b2cf-77ce-484a-85b8-fd161500be1e",
		Type: "style",
		Name: "Стиль",
	}
)

// NewsType <struct>
// is used to describe news type model.
type NewsType struct {
	ID   string `json:"_id" bson:"_id"`
	Type string `json:"type" bson:"type"`
	Name string `json:"name" bson:"name"`
}

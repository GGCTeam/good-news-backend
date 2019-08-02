package models

const (
	TypeNews     string = "news"
	TypePost     string = "post"
	TypeBusiness string = "business"
	TypeStyle    string = "style"
)

// NewsType <struct>
// is used to describe news type model.
type NewsType struct {
	Type string `json:"type" bson:"type"`
	Name string `json:"name" bson:"name"`
}

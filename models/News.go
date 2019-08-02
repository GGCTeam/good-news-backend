package models

// News <model>
// is used to describe article model.
type News struct {
	Type       string     `json:"type" bson:"type"`
	Title      string     `json:"title" bson:"title"`
	Preamble   string     `json:"preamble" bson:"preamble"`
	TimeAdded  uint64     `json:"time_added" bson:"time_added"`
	Link       string     `json:"link" bson:"link"`
	NewsType   NewsType   `json:"news_type" bson:"news_type"`
	NewsSource NewsSource `json:"news_source" bson:"news_source"`
}

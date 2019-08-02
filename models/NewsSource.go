package models

// NewsSource <model>
// is used to describe news source model.
type NewsSource struct {
	Name string `json:"name" bson:"name"`
	// Types
	ImageURL string `json:"image_url" bson:"image_url"`
}

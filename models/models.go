package models

type Article struct {
	ID      string `bson:"_id"` //primary id: _id (mongodb)
	Title   string `json:"title,omitempty" bson:"title,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
}

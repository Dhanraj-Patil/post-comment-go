package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Post struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	PostData string        `gorm:"not null" json:"post_data" bson:"post_data"`
	Thread   string        `json:"thread" bson:"thread"`
	CreateAt time.Time
}

type Thread struct {
	Post     Post   `json:"post"`
	Comments []Post `json:"comments"`
}

type PostRequest struct {
	PostData string `gorm:"not null" json:"post_data" bson:"post_data"`
	Thread   string `json:"thread" bson:"thread"`
}

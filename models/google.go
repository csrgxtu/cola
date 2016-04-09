package models

type (
  GoogleResult struct {
    Title string `json:"title" bson:"title"`
    LongUrl string `json:"longurl" bson:"longurl"`
    ShortUrl string `json:"shorturl" bson:"shorturl"`
    Abstract string `json:"abstract" bson:"abstract"`
  }
)

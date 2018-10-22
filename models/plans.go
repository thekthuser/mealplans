package models

import (
  "gopkg.in/mgo.v2/bson"
)

type Plan struct {
  Id bson.ObjectId `bson:"_id" json:"id"`
  Name string `bson:"name" json:"name"`
  Cost int `bson:"cost" json:"cost"`
  Market string `bson:"market" json:"market"`
  Semester1Start string `bson:"semester1start" json:"semester1start"`
  Semester1End string `bson:"semester1end" json:"semester1end"`
  Semester2Start string `bson:"semester2start" json:"semester2start"`
  Semester2End string `bson:"semester2end" json:"semester2end"`
  Semester3Start string `bson:"semester3start" json:"semester3start"`
  Semester3End string `bson:"semester3end" json:"semester3end"`
  MarketingText1 string `bson:"marketingtext1" json:"marketingtext1"`
  MarketingText2 string `bson:"marketingtext2" json:"marketingtext2"`
  MarketingText3 string `bson:"marketingtext3" json:"marketingtext3"`
}

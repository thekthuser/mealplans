package models

import (
  "gopkg.in/mgo.v2/bson"
)

type Plan struct {
  Id bson.ObjectId `bson:"_id" json:"id"`
  Name string `bson:"name" json:"name"`
  Cost string `bson:"cost" json:"cost"`
  Market string `bson:"market" json:"market"`
  //Semester1Start string `bson:"semester1start" json:"semester1start"`
  MarketingText1 string `bson:"marketingtext1" json:"marketingtext1"`
  MarketingText2 string `bson:"marketingtext2" json:"marketingtext2"`
  MarketingText3 string `bson:"marketingtext3" json:"marketingtext3"`
}

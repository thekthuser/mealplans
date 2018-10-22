package models

import (
  "gopkg.in/mgo.v2/bson"
)

type User struct {
  id bson.ObjectId `bson:"_id" json:"id"`
  mealPlanId string `bson:"mealPlanId" json:"mealPlanId"`
  username string `bson:"username" json:"username"`
  password string `bson:"password" json:"password"`
  name string `bson:"name" json:"name"`
  school string `bson:"school" json:"school"`
}

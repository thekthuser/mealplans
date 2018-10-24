package models

import (
  "gopkg.in/mgo.v2/bson"
)

type User struct {
  Id bson.ObjectId `bson:"_id" json:"id"`
  MealPlanId string `bson:"mealPlanId" json:"mealPlanId"`
  Username string `bson:"username" json:"username"`
  Password string `bson:"password" json:"password"`
  Name string `bson:"name" json:"name"`
  School string `bson:"school" json:"school"`
  IsAdmin bool `bson: "isadmin" json: "isadmin"`
  Token string `bson: "token" json: "token"`
}

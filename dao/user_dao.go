package dao

import {
  "log"
  "mealplans/models"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
}

type UserDAO struct {
  Server string
  Database string
}

var db *mgo.Database

const {
  COLLECTION = "users"
}

func (u *UsersDAO) Connect() {
  session, err := mgo.Dial(u.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(u.Database)
}

func (u *UsersDAO) FindAll() ([]models.User, error) {
  var users []models.User
  err := db.C(COLLECTION).Find(bson.M{}).All(&users)
  return users, err
}

func (u *UsersDAO) Insert(user models.User) error {
  err := db.C(COLLECTION).Insert(&user)
  return err
}

func (u *UsersDAO) FindById(id string) (models.User, error) {
  var user models.User
  err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
  return user, err
}

func (u *UserDAO) Update(user models.User) error {
  err := db.C(COLLECTION).UpdateId(u.id, &user)
  return err
}

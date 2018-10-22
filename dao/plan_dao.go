package dao

import (
  "log"
  "mealplans/models"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type PlanDAO struct {
  Server string
  Database string
}

var db *mgo.Database

const (
  PLAN_COLLECTION = "plans"
)

func (p *PlanDAO) Connect() {
  session, err := mgo.Dial(p.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(p.Database)
}

func (p *PlanDAO) FindAll() ([]models.Plan, error) {
  var plans []models.Plan
  err := db.C(PLAN_COLLECTION).Find(bson.M{}).All(&plans)
  return plans, err
}

func (p *PlanDAO) Insert(plan models.Plan) error {
  err := db.C(PLAN_COLLECTION).Insert(&plan)
  return err
}

func (p *PlanDAO) FindById(id string) (models.Plan, error) {
  var plan models.Plan
  err := db.C(PLAN_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&plan)
  return plan, err
}

/*
func (u *UserDAO) Update(user models.User) error {
  err := db.C(COLLECTION).UpdateId(u.id, &user)
  return err
}
*/

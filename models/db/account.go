package db 

import (
	"time"
	"labix.org/v2/mgo"
        "labix.org/v2/mgo/bson"
)

type Account struct {
        Name string
        Email string
        Phone string
        Password string
        Type string
        Date time.Time
}

func (a Account) retrieve(name string,s *mgo.Collection) Account{
        result := Account{}
	err := s.Find(bson.M{"name": name}).One(&result)
        if err != nil {
                panic(err)
        }
        return result
}

func insert(u Account,s *mgo.Collection){
	err := s.Insert(&u)
        if err != nil {
                panic(err)
        }
}

func remove(u Account,s *mgo.Collection){
	err := s.Remove(&u)
        if err != nil {
                panic(err)
        }
}

func update(u Account,s *mgo.Collection){
	colQuerier := bson.M{"name": u.Name}
	change := bson.M{"$set": bson.M{"email": u.Email,"phone": u.Phone,"password": u.Password}}
	err := s.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
}

func (a Account) GetAccount(name string) Account{
	session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("user")	
	b := Account{}
	b = b.retrieve(name, c)
	return b
}


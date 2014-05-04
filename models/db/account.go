package db 

import (
	"time"
        "fmt"
        "labix.org/v2/mgo"
        "labix.org/v2/mgo/bson"
)

type User struct {
        Name string
        Email string
        Phone string
        Password string
        Type string
        dataCreated time.Time
}

func retrieve(name string,s *mgo.Collection) User{
        result := User{}
	err := s.Find(bson.M{"name": name}).One(&result)
        if err != nil {
                panic(err)
        }
        return result
}

func insert(u User,s *mgo.Collection){
	err := s.Insert(&u)
        if err != nil {
                panic(err)
        }
}

func remove(u User,s *mgo.Collection){
	err := s.Remove(&u)
        if err != nil {
                panic(err)
        }
}

func update(u User,s *mgo.Collection){
	colQuerier := bson.M{"name": u.Name}
	change := bson.M{"$set": bson.M{"email": u.Email,"phone": u.Phone,"password": u.Password}}
	err := s.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
}


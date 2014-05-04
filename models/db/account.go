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

// Retrieve Account into Mongo
func (a Account) retrieve(name string,s *mgo.Collection) Account{
        result := Account{}
	err := s.Find(bson.M{"name": name}).One(&result)
        if err != nil {
                panic(err)
        }
        return result
}

// Insert Account into Mongo
func insert(u Account,s *mgo.Collection){
	err := s.Insert(&u)
        if err != nil {
                panic(err)
        }
}

// Remove Account into Mongo
func remove(u Account,s *mgo.Collection){
	err := s.Remove(&u)
        if err != nil {
                panic(err)
        }
}

// Update Account into Mongo
func update(u Account,s *mgo.Collection){
	colQuerier := bson.M{"name": u.Name}
	change := bson.M{"$set": bson.M{"email": u.Email,"phone": u.Phone,"password": u.Password, "type": u.Type}}
	err := s.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
}

// Retreive Account into Mongo
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

// Additional Account into Mongo
func (a Account) AddAccount(){
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("test").C("user")
        insert(a, c)
}

// Update Account into Mongo
func (a Account) UpdateAccount(){
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("test").C("user")
        update(a, c)
}

// Delete Account into Mongo
func (a Account) DeleteAccount(){
	session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("test").C("user")
        remove(a, c)
}

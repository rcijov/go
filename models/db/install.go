package main

import (
        "fmt"
	"time"
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

type Job struct{
	UID string
	title string
	job_description string
	paid int
	when string
	workHour int
	people_wanted int
}

func main() {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("user")
        err = c.Insert(&User{"dada", "da@da.com", "41443", "da", "admin", time.Now()},
	               &User{"lala", "ca@ca.com", "21423", "ca", "user", time.Now()})
        if err != nil {
                panic(err)
        }

        result := User{}
        err = c.Find(bson.M{"name": "dada"}).One(&result)
        if err != nil {
                panic(err)
        }

        fmt.Println("Phone:", result.Phone)
}

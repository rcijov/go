package models

import (
	"web/models/db"
	"web/models/entities"
	"time"
)

type Account struct {
    	Name string
    	Email string
    	Phone string
    	Password string
    	Date time.Time
}

func (b Account) Account() Account{	
	acc := Account{}
	a := entities.Account{}	
 	a.SetAcc("a","b","c","d", time.Now())

	acc.Name = a.GetName()
	acc.Email = a.GetEmail()
	acc.Phone = a.GetPhone()
	acc.Password = a.GetPassword()
	acc.Date = a.GetDateCreated()
 
	acc = GetAccount("lala")

	return acc
}

func GetAccount(name string) Account{

	b := db.Account{}
	a := db.Account{}
	b = a.GetAccount(name)
	c := Account{}
	c.Name = b.Name
	c.Email = b.Email
	c.Phone = b.Phone
	c.Password = b.Password
	c.Date = b.Date
	return c
}

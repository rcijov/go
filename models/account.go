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
	Type string
    	Date time.Time
}

// Account Main Function
func (b Account) Account() Account{	
	acc := Account{}	
	acc = GetAccount("dada")
	return acc
}

// Modify Account - old and new changes
func ModifyAccount(oldAcc *entities.Account, newAcc entities.Account){
	if newAcc.Email != ""{
		oldAcc.Email = newAcc.Email
	}
	if newAcc.Phone != ""{
		oldAcc.Phone = newAcc.Phone
	}
	if newAcc.Type != ""{
		oldAcc.Type = newAcc.Type
	}
	if newAcc.Password != ""{
	 oldAcc.Password = newAcc.Password
	}
}

// Account to Entities Account
func AccToEnt(account Account) entities.Account{
	a:= entities.Account{}
	a.Name = account.Name
	a.Email = account.Email
	a.Phone = account.Phone
	a.Password = account.Password
	a.Type = account.Type
	a.Date = account.Date
	return a
}

// Additional Account
func AddAccount(acc entities.Account){
	a := db.Account{acc.Name,acc.Email,acc.Phone,acc.Password,acc.Type,acc.Date}
	a.AddAccount()
}

// Update Account
func UpdAccount(acc entities.Account){
	a := db.Account{acc.Name,acc.Email,acc.Phone,acc.Password,acc.Type,acc.Date}
	a.UpdateAccount()
}

// Delete Account
func DelAccount(acc entities.Account){
        a := db.Account{acc.Name,acc.Email,acc.Phone,acc.Password,acc.Type,acc.Date}
        a.DeleteAccount()
}

// Retrieve Account
func GetAccount(name string) Account{
	b := db.Account{}
	a := db.Account{}
	b = a.GetAccount(name)
	c := Account{}
	c.Name = b.Name
	c.Email = b.Email
	c.Phone = b.Phone
	c.Type = b.Type
	c.Password = b.Password
	c.Date = b.Date
	return c
}

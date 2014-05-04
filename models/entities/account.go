package entities

import (
	"time"
)

type Account struct {
    name string
    email string
    phone string
    password string
    dateCreated time.Time
}

// Set Account
func (f *Account) SetAcc(name string, email string, phone string, password string, datecreated time.Time) {
    f.name = name
    f.email = email
    f.phone = phone
    f.password = password
    f.dateCreated = datecreated
}

// SetName Account
func (f *Account) SetName(name string) {
    f.name = name
}

// GetName Account
func (f Account) GetName() string {
    return f.name
}

// SetEmail Account
func (f *Account) SetEmail(email string){
    f.email = email
}

// GetEmail Account
func (f Account) GetEmail() string {
    return f.email
}

// SetPhone Account
func (f *Account) SetPhone(phone string){
    f.phone = phone
}

// GetPhone Account
func (f Account) GetPhone() string {
    return f.phone
}

// SetPassword Account
func (f *Account) SetPassword(password string){
    f.password = password
}

// GetPassword Account
func (f Account) GetPassword() string {
    return f.password
}

// SetDateCreated Account
func (f *Account) SetDateCreated(date time.Time){  
     f.dateCreated = date
}

// GetDateCreated Account
func (f Account) GetDateCreated() time.Time {
    return f.dateCreated
}


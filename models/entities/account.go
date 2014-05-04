package entities

import (
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

// Set Account
func (f *Account) SetAcc(name string, email string, phone string, password string, typez string, datecreated time.Time) {
    f.Name = name
    f.Email = email
    f.Phone = phone
    f.Password = password
    f.Type = typez
    f.Date = datecreated
}

// SetName Account
func (f *Account) SetName(name string) {
    f.Name = name
}

// GetName Account
func (f Account) GetName() string {
    return f.Name
}

// SetEmail Account
func (f *Account) SetEmail(email string){
    f.Email = email
}

// GetEmail Account
func (f Account) GetEmail() string {
    return f.Email
}

// SetPhone Account
func (f *Account) SetPhone(phone string){
    f.Phone = phone
}

// GetPhone Account
func (f Account) GetPhone() string {
    return f.Phone
}

// SetPassword Account
func (f *Account) SetPassword(password string){
    f.Password = password
}

// GetPassword Account
func (f Account) GetPassword() string {
    return f.Password
}

// SetDateCreated Account
func (f *Account) SetDateCreated(date time.Time){  
     f.Date = date
}

// GetDateCreated Account
func (f Account) GetDateCreated() time.Time {
    return f.Date
}

// SetType Account
func (f *Account) SetType(typez string){
     f.Type = typez
}

// GetType Account
func (f Account) GetType() string{
    return f.Type
}

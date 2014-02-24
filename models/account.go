package models

type Account struct {
    name string
    email string
    phone string
    password string
    dateCreated string
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

// SetdateCreated Account
func (f *Account) SetDateCreated(dateCreated string){
    f.dateCreated = dateCreated
}

// GetPassword Account
func (f Account) GetDateCreated() string {
    return f.dateCreated
}


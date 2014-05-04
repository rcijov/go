package models

type Job struct {
    title string
    description string
    paid int
    date string
    workHours int
    peopleWanted int
}

// SetTitle Job
func (f *Job) SetTitle(title string) {
    f.title = title
}

// GetTitle Job
func (f Job) GetTitle() string {
    return f.title
}

// SetDescription Job    
func (f *Job) SetDescription(description string) {
    f.description = description
}

// GetDescription Job    
func (f Job) GetDescription() string {
    return f.description
}

// SetPaid Job
func (f *Job) SetPaid(paid int) {
    f.paid = paid
}

// GetPaid Job
func (f Job) GetPaid() int {
    return f.paid
}

// SetDate Job
func (f *Job) SetDate(date string) {
    f.date = date
}

// GetDate Job
func (f Job) GetDate() string {
    return f.date
}

// SetWorkHours Job
func (f *Job) SetWorkHours(workHours int) {
    f.workHours = workHours
}

// GetWorkHours Job
func (f Job) GetWorkHours() int {
    return f.workHours 
}

// SetPeopleWanted Job
func (f *Job) SetPeopleWanted(peopleWanted int) {
    f.peopleWanted = peopleWanted
}

// GetPeopleWanted Job
func (f Job) GetPeopleWanted() int {
    return f.peopleWanted
}





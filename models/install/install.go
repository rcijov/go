package install

import (
       
        "labix.org/v2/mgo"
        "time"
        "labix.org/v2/mgo/bson"
)

type User_Profile struct {
		ID int
        Name string
        Email string
        Phone string
        password string
		dataCreated time.Time 
        
}

type Job struct{
	userID int
	ID int
	title string
	job_description string
	paid int
	when string
	workHour int
	people_wanted int
}




func DatabaseInit() {
        mongoSession, err := mgo.Dial("localhost"); // connect to database
        if err != nil {
        		panic(err)
    	}
    defer mongoSession.Close()
     mongoSession.SetMode(mgo.Monotonic, true)
     
   database  := mongoSession.DB("mc_hack")
   account := database.C("Account")
  	job := database.C("Job")
   
   
    
       err = account.Insert(&User_Profile{1, "Mohammed" , "m.f.e@live.ca" , "514 497 7551", "1234" , time.Now() })

        if err != nil {
        	panic(err)
        }
        result := User_Profile{}
        err = account.Find(bson.M{"Name" : "Mohammed"}).One(&result)
        if err != nil {
        	panic(err)
        }
     err = job.Insert(&Job{result.ID, 1 , "Google" , "cleanner", 34 , "Saturday ", 5 , 2 })
   if err != nil {
   		panic(err)
   	}
   	return 
   	
    
}

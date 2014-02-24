package controllers


type Home struct {
	flag int
}

func (a Home) Home() int {
	return a.flag	
}

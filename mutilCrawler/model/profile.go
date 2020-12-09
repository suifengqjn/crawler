package model

import "fmt"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

func (p *Profile)String() string  {
	re := p.Name
	re += p.Gender
	re += fmt.Sprint("age :", p.Age)
	re += fmt.Sprint("height :", p.Height)
	re += fmt.Sprint("weight :", p.Weight)
	re += p.Income
	re += p.Marriage
	re += p.Education
	re += p.Occupation
	re += p.Hokou
	re += p.Xinzuo
	re += p.House
	re += p.Car + `\n`
	return re
}
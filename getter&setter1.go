package main

import "fmt"

type person struct{
	name string
	age int
	job string
}

func (p person) GetName()string{
	return p.name
}
func (s *person) SetName(name string) {
	s.name = "shovon"
}

func GetterSetter1 () {
	person1 := person{
         name: "pranto",
		 age: 24,
		 job: "Intern software engineer", 
	}
    fmt.Println("the name is",person1.GetName())
	person1.SetName(person1.name)
	fmt.Println("Name is : ",person1.GetName())
}
package main

import "fmt"

type student struct {
	Name string
	Roll int
	subject
}
type subject struct {
	sub   string
	marks int
}

func (p student) GetValue() string {
	return p.sub
}
func (q *subject) SetValue(sub string) {
    q.sub = "ICT"
}


func GetterSetter2() {
	per1 := student{
		Name: "Pranto",
		Roll: 10,
		subject: subject{
			sub:   "Bangla",
			marks: 70,
		},
	}
	fmt.Println("Subject ", per1.GetValue())
	per1.SetValue(per1.sub)
	fmt.Println("Subject ",per1.GetValue())

}

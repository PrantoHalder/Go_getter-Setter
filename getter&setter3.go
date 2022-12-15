package main

import "fmt"

type sub struct{
	sub string
	mark int
	subjectCode int
}
func (p sub) GetValue()string{
	return p.sub
}
func (q *sub) SetValue(sub string){
	q.sub = "ICT"
}
func GetterSetter3(){
  per2 := sub{
  	sub:         "Bangla",
  	mark:        75,
  	subjectCode: 110,
  }
  fmt.Println("subject :",per2.GetValue())
  per2.SetValue(per2.sub)
  fmt.Println("Changed Subject ",per2.GetValue())

}
package jobs

import "fmt"

type Person struct{
	name string
}

func NewPerson(name string)*Person{
	return &Person{
		name: name,
	}
}
func (p *Person)Run()  {
	fmt.Println(p.name)
}
package pattern

import (
	"fmt"
	"testing"
)

type Person struct {
	Name, Sexual string
	Age          int
}

func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func Test_Person(t *testing.T) {
	var p = Person{
		Name:   "Hao Chen",
		Sexual: "Male",
		Age:    44,
	}

	PrintPerson(&p)
	p.Print()
}

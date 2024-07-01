package main

import (
	"fmt"
)

type Project struct {
	value int
}

func (p *Project) Hello() {
	p.value += 5
	fmt.Println(p.value)
}

func (p *Project) Hey() {
	p.value += 10
	fmt.Println(p.value)
}

func main() {
	var p *Project
	p.Hello()
	p.Hey()
}

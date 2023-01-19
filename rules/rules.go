package rules

import "fmt"

type Rules struct{}

type Interface interface {
	Rulesfunc()
}

func (r *Rules) Rulesfunc() {
	fmt.Println("this is rules function")
}
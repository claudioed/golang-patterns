package main

import (
	"fmt"
	"golang-patterns/chain-of-responsibility/chain"
	"golang-patterns/chain-of-responsibility/chain/rules"
	"golang-patterns/chain-of-responsibility/chain/rules/domain"
)

func main() {
	p := &chain.Person{
		Age:     19,
		Gender:  "female",
		Country: "BR",
	}
	cc := &domain.ClassificationContainer{
		Person: p,
	}

	genderRule := &rules.GenderRule{Gender: "male"}
	ageRule := &rules.AgeGreaterThan{
		Age:      18,
		NextStep: genderRule,
	}
	ageRule.Execute(cc)
	fmt.Println(cc.Results)
}

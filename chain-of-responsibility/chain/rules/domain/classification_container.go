package domain

import "golang-patterns/chain-of-responsibility/chain"

type ClassificationContainer struct {
	Person  *chain.Person
	Results []*Result
}

func (cc *ClassificationContainer) AddResult(result *Result) {
	cc.Results = append(cc.Results, result)
}

package rules

import (
	"golang-patterns/chain-of-responsibility/chain/rules/domain"
)

type AgeGreaterThan struct {
	Age      int
	NextStep ClassificationRule
}

func (agt *AgeGreaterThan) Execute(container *domain.ClassificationContainer) {
	if container.Person.Age < agt.Age {
		container.AddResult(ageResultFailed(container))
	}
	if agt.NextStep != nil {
		agt.NextStep.Execute(container)
	}
}
func (agt *AgeGreaterThan) next(next ClassificationRule) {
	agt.NextStep = next
}

func ageResultFailed(container *domain.ClassificationContainer) *domain.Result {
	return &domain.Result{
		Rule:    "age",
		Success: false,
		Reason:  "Age less than allowed",
	}
}

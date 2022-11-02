package rules

import "golang-patterns/chain-of-responsibility/chain/rules/domain"

type GenderRule struct {
	Gender   string
	NextStep ClassificationRule
}

func (gr *GenderRule) Execute(container *domain.ClassificationContainer) {
	if gr.Gender != container.Person.Gender {
		container.AddResult(genderResultFailed(container))
	}
	if gr.NextStep != nil {
		gr.NextStep.Execute(container)
	}
}
func (gr *GenderRule) next(next ClassificationRule) {
	gr.NextStep = next
}

func genderResultFailed(container *domain.ClassificationContainer) *domain.Result {
	return &domain.Result{
		Rule:    "gender",
		Success: false,
		Reason:  "Gender not allowed",
	}
}

package rules

import "golang-patterns/chain-of-responsibility/chain/rules/domain"

type ClassificationRule interface {
	Execute(container *domain.ClassificationContainer)
	next(next ClassificationRule)
}

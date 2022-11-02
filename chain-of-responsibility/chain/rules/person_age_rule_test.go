package rules

import (
	"github.com/stretchr/testify/assert"
	"golang-patterns/chain-of-responsibility/chain"
	"golang-patterns/chain-of-responsibility/chain/rules/domain"
	"testing"
)

func TestCheckAgeGreaterThan(t *testing.T) {
	p := &chain.Person{
		Age:     19,
		Gender:  "male",
		Country: "BR",
	}
	cc := &domain.ClassificationContainer{
		Person: p,
	}
	rule := &AgeGreaterThan{
		Age: 18,
	}
	rule.Execute(cc)
	assert.Equal(t, 0, len(cc.Results), "Result should be empty")
}

func TestCheckAgeLessThan(t *testing.T) {
	p := &chain.Person{
		Age:     16,
		Gender:  "female",
		Country: "BR",
	}
	cc := &domain.ClassificationContainer{
		Person: p,
	}
	rule := &AgeGreaterThan{
		Age: 18,
	}
	rule.Execute(cc)
	assert.Equal(t, 1, len(cc.Results), "Result should be not empty")
}

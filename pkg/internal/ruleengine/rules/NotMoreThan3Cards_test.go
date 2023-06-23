package rules

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotMoreThan3Cards_testRule(t *testing.T) {
	tests := []struct {
		Name           string
		TestInput      RuleParams
		ExpectedResult bool
		ExpectedError  error
	}{
		{
			Name: "Number of cards is less than 3",
			TestInput: RuleParams{
				NumberOfCreditCards: 2,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Number of cards is more than 3",
			TestInput: RuleParams{
				NumberOfCreditCards: 4,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Number of cards are 3",
			TestInput: RuleParams{
				NumberOfCreditCards: 3,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Number of cards is less than 0",
			TestInput: RuleParams{
				NumberOfCreditCards: -1,
			},
			ExpectedResult: false,
			ExpectedError:  errors.New("number of cards can't be negative"),
		},
		{
			Name: "Number of cards are 0",
			TestInput: RuleParams{
				NumberOfCreditCards: 0,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
	}

	for _, test := range tests {
		notMoreThan3CardsRule := NotMoreThan3Cards{}
		result, err := notMoreThan3CardsRule.Evaluate(test.TestInput)
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)

	}
}

package rules

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowCreditRiskScore(t *testing.T) {
	tests := []struct {
		Name           string
		TestInput      RuleParams
		ExpectedResult bool
		ExpectedError  error
	}{
		{
			Name: "Age is more than 18 and numbers of cards are 2",
			TestInput: RuleParams{
				Age:                 19,
				NumberOfCreditCards: 2,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Age is less than 18 and numbers of cards are 2",
			TestInput: RuleParams{
				Age:                 17,
				NumberOfCreditCards: 2,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Age is equal than 18 and numbers of cards are 2",
			TestInput: RuleParams{
				Age:                 18,
				NumberOfCreditCards: 2,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Age is more than 18 and numbers of cards are 3",
			TestInput: RuleParams{
				Age:                 19,
				NumberOfCreditCards: 3,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Age is less than 18 and numbers of cards are 3",
			TestInput: RuleParams{
				Age:                 17,
				NumberOfCreditCards: 3,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Age is equal than 18 and numbers of cards are 3",
			TestInput: RuleParams{
				Age:                 18,
				NumberOfCreditCards: 3,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Age is less than 0 and numbers of cards are 3",
			TestInput: RuleParams{
				Age:                 -1,
				NumberOfCreditCards: 3,
			},
			ExpectedResult: false,
			ExpectedError:  errors.New("age and numbers of cards should be non negative"),
		},
		{
			Name: "Age is less than 0 and numbers of cards are less than 0",
			TestInput: RuleParams{
				Age:                 -1,
				NumberOfCreditCards: -3,
			},
			ExpectedResult: false,
			ExpectedError:  errors.New("age and numbers of cards should be non negative"),
		},
		{
			Name: "Age is 18 than 0 and numbers of cards are 3",
			TestInput: RuleParams{
				Age:                 18,
				NumberOfCreditCards: -1,
			},
			ExpectedResult: false,
			ExpectedError:  errors.New("age and numbers of cards should be non negative"),
		},
	}

	for _, test := range tests {
		lowCreditRiskScoreRule := LowCreditRiskScore{}
		result, err := lowCreditRiskScoreRule.Evaluate(test.TestInput)
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)

	}
}

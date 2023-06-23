package rules

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEarnMoreThan100000(t *testing.T) {
	tests := []struct {
		Name           string
		TestInput      RuleParams
		ExpectedResult bool
		ExpectedError  error
	}{
		{
			Name: "Income is less than 100000",
			TestInput: RuleParams{
				Income: 45433,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Income is more than 100000",
			TestInput: RuleParams{
				Income: 238600,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Income is equal to 100000",
			TestInput: RuleParams{
				Income: 100000,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Income is negative",
			TestInput: RuleParams{
				Income: -100,
			},
			ExpectedResult: false,
			ExpectedError:  errors.New("income is negative"),
		},
		{
			Name: "Income is zero",
			TestInput: RuleParams{
				Income: 0,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
	}

	for _, test := range tests {
		earnMoreThan100000Rule := EarnMoreThan100000{}
		result, err := earnMoreThan100000Rule.Evaluate(test.TestInput)
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)

	}
}

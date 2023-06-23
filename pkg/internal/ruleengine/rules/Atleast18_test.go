package rules

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtleast18Rule(t *testing.T) {
	tests := []struct {
		Name           string
		TestInput      RuleParams
		ExpectedResult bool
		ExpectedError  error
	}{
		{
			Name: "Age is less than 18",
			TestInput: RuleParams{
				Age: 17,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Age is more than 18",
			TestInput: RuleParams{
				Age: 19,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Age is equal to 18",
			TestInput: RuleParams{
				Age: 18,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Age is negative",
			TestInput: RuleParams{
				Age: -1,
			},
			ExpectedResult: false,
			ExpectedError:  errors.New("age is negative"),
		},
		{
			Name: "Age is zero",
			TestInput: RuleParams{
				Age: 0,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
	}

	for _, test := range tests {
		atLeast18rule := Atleast18{}
		result, err := atLeast18rule.Evaluate(test.TestInput)
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)
	}
}

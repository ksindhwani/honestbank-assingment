package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotPoliticallyExposedRule(t *testing.T) {
	tests := []struct {
		Name           string
		TestInput      RuleParams
		ExpectedResult bool
		ExpectedError  error
	}{
		{
			Name: "Is Politically Exposed",
			TestInput: RuleParams{
				PoliticallyExposed: true,
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Is Not Politically Exposed",
			TestInput: RuleParams{
				PoliticallyExposed: false,
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
	}

	for _, test := range tests {
		notPoliticallyExposedRule := NotPoliticallyExposed{}
		result, err := notPoliticallyExposedRule.Evaluate(test.TestInput)
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)

	}
}

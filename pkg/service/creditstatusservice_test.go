package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/honestbank/tech-assignment-backend-engineer/pkg/internal/ruleengine/rules"
	"github.com/stretchr/testify/assert"
)

func TestProcessCreditCardApplication(t *testing.T) {
	tests := []struct {
		Name           string
		Input          RecordData
		ExpectedResult JSONResponse
		ExpectedError  error
	}{
		{
			Name: "All parmas are good and valid",
			Input: RecordData{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Approved},
			ExpectedError:  nil,
		},
		{
			Name: "Age is less than 18",
			Input: RecordData{
				Income:              239700,
				Age:                 18,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Declined},
			ExpectedError:  nil,
		},
		{
			Name: "Age is negative",
			Input: RecordData{
				Income:              239700,
				Age:                 -18,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{},
			ExpectedError:  fmt.Errorf("Error occured while executing rule. %w", fmt.Errorf("Error: %w", errors.New("age is negative"))),
		},
		{
			Name: "Income is less than 100000",
			Input: RecordData{
				Income:              34560,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Declined},
			ExpectedError:  nil,
		},
		{
			Name: "Income is negative",
			Input: RecordData{
				Income:              -100,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{},
			ExpectedError:  fmt.Errorf("Error occured while executing rule. %w", fmt.Errorf("Error: %w", errors.New("income is negative"))),
		},
		{
			Name: "Number of credit cards are 4 ",
			Input: RecordData{
				Income:              234500,
				Age:                 20,
				NumberOfCreditCards: 4,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Declined},
			ExpectedError:  nil,
		},
		{
			Name: "Number of credit card is negative",
			Input: RecordData{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: -1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{},
			ExpectedError:  fmt.Errorf("Error occured while executing rule. %w", fmt.Errorf("Error: %w", errors.New("number of cards can't be negative"))),
		},
		{
			Name: "Is Politically Exposed",
			Input: RecordData{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  true,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Declined},
			ExpectedError:  nil,
		},
		{
			Name: "Phone number is in area code 8",
			Input: RecordData{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "876-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Approved},
			ExpectedError:  nil,
		},
		{
			Name: "Phone number is in invalid area code",
			Input: RecordData{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "-576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{},
			ExpectedError:  fmt.Errorf("Error occured while executing rule. %w", fmt.Errorf("Error: %w", errors.New("area code should be in between 0-9"))),
		},
		{
			Name: "Age is less than 18 and politically exposed",
			Input: RecordData{
				Income:              239700,
				Age:                 17,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  true,
				PhoneNumber:         "576-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Declined},
			ExpectedError:  nil,
		},
		{
			Name: "Credit Score is medium",
			Input: RecordData{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 2,
				PoliticallyExposed:  false,
				PhoneNumber:         "476-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Declined},
			ExpectedError:  nil,
		},
		{
			Name: "Phone number is in pre-approved list, Credit Score is medium",
			Input: RecordData{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 2,
				PoliticallyExposed:  false,
				PhoneNumber:         "323-912-6988",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Approved},
			ExpectedError:  nil,
		},
		{
			Name: "Phone number is not in pre-approved list, Credit Score is medium",
			Input: RecordData{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 2,
				PoliticallyExposed:  false,
				PhoneNumber:         "416-376-123",
				JobIndustryCode:     "2-317 - Select Borrow",
			},
			ExpectedResult: JSONResponse{Status: Declined},
			ExpectedError:  nil,
		},
	}

	for _, test := range tests {
		result, err := ProcessCreditCardApplication(test.Input)
		rules.APPROVED_PHONE_NUMBERS_FILE_NAME = "pre-approved-phone-numbers_test.txt"
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)

	}
}

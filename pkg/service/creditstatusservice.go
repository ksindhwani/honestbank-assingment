package service

import (
	"fmt"

	"github.com/honestbank/tech-assignment-backend-engineer/pkg/internal/ruleengine"
	"github.com/honestbank/tech-assignment-backend-engineer/pkg/internal/ruleengine/rules"
)

const (
	Approved = "approved"
	Declined = "declined"
)

type RecordData struct {
	Income              int    `json:"income"`
	NumberOfCreditCards int    `json:"number_of_credit_cards"`
	Age                 int    `json:"age"`
	PoliticallyExposed  bool   `json:"politically_exposed"`
	JobIndustryCode     string `json:"job_industry_code"`
	PhoneNumber         string `json:"phone_number"`
}

type JSONResponse struct {
	Status string `json:"status"`
}

func ProcessCreditCardApplication(data RecordData) (JSONResponse, error) {
	ruleEngine := ruleengine.PrepareNewRuleEngine()
	ok, err := ruleEngine.ValidateCreditCardApplication(prepareRuleParams(data))
	if err != nil {
		return JSONResponse{}, fmt.Errorf("Error occured while executing rule. %w", err)
	}
	if !ok {
		return JSONResponse{Status: Declined}, nil
	}

	return JSONResponse{Status: Approved}, nil
}

func prepareRuleParams(data RecordData) rules.RuleParams {
	return rules.RuleParams{
		Age:                 data.Age,
		Income:              data.Income,
		NumberOfCreditCards: data.NumberOfCreditCards,
		PoliticallyExposed:  data.PoliticallyExposed,
		PhoneNumber:         data.PhoneNumber,
	}
}

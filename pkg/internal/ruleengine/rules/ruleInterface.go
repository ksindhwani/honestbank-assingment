package rules

type RuleParams struct {
	Income              int
	NumberOfCreditCards int
	Age                 int
	PoliticallyExposed  bool
	PhoneNumber         string
}

type Rule interface {
	Evaluate(Rule RuleParams) (bool, error)
}

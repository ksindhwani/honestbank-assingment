package rules

type NotPoliticallyExposed struct{}

func (r NotPoliticallyExposed) Evaluate(params RuleParams) (bool, error) {
	return !params.PoliticallyExposed, nil
}

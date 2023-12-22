package provider

func (rp *ruleSetProvider) Count() int {
	return rp.strategy.Count()
}

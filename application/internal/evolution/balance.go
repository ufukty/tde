package evolution

type Balance struct {
	EndedSearches int
}

type BalanceBook []Balance

func (b *BalanceBook) NewGen() {
	(*b) = append((*b), Balance{})
}

func (b *BalanceBook) Current() *Balance {
	return &(*b)[len(*b)-1]
}

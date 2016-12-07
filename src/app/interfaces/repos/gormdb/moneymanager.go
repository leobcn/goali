package gormdb

import "app"

func NewMoneyManager(r *Repo) *MoneyManager {
	return &MoneyManager{r}
}

type MoneyManager struct {
	*Repo
}

func (mm *MoneyManager) GetTransactions(w app.DBWhere, f *app.DBFilter) ([]app.Transaction, error) {
	var es []app.Transaction

	return es, mm.FindBy(&es, w, f)
}

func (mm *MoneyManager) GetAccounts(w app.DBWhere, f *app.DBFilter) ([]app.Account, error) {
	var as []app.Account

	return as, mm.FindBy(&as, app.DBWhere{}, f)
}

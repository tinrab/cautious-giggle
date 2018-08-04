package domain

type Account struct {
	Entity
	Name string
}

func NewAccount(name string) Account {
	return Account{
		Entity: newEntity(),
		Name:   name,
	}
}

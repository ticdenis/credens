package domain

type (
	AccountBuilder interface {
		Build(
			id string,
			name string,
			username string,
			password string,
		) (*Account, error)
	}

	BaseAccountBuilder struct {
	}
)

func NewAccountBuilder() *BaseAccountBuilder {
	return &BaseAccountBuilder{}
}

func (_ BaseAccountBuilder) Build(
	id string,
	name string,
	username string,
	password string,
) (*Account, error) {
	return &Account{
		id:       NewAccountId(id),
		name:     NewAccountName(name),
		username: NewAccountUsername(username),
		password: NewAccountPassword(password),
	}, nil
}

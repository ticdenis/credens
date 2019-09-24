package read

import (
	"credens/libs/accounts/domain"
)

type ReadAccountResponse struct {
	Id       string
	Name     string
	Username string
}

func NewReadAccountResponse(account domain.Account) *ReadAccountResponse {
	return &ReadAccountResponse{
		account.Id().Value(),
		account.Name().Value(),
		account.Username().Value(),
	}
}

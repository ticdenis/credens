package read

import "credens/src/domain/account"

type ReadAccountResponse struct {
	Id       string
	Name     string
	Username string
}

func NewReadAccountResponse(account account.Account) *ReadAccountResponse {
	return &ReadAccountResponse{
		account.Id().Value(),
		account.Name().Value(),
		account.Username().Value(),
	}
}

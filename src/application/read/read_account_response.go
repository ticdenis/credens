package read

import "credens/src/domain/account"

type ReadAccountResponse struct {
	Id       string `json:"Id"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
}

func NewReadAccountResponse(account account.Account) *ReadAccountResponse {
	return &ReadAccountResponse{
		account.Id().Value(),
		account.Name().Value(),
		account.Username().Value(),
	}
}

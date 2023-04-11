package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetAllUserTransactions struct {
	XMLName xml.Name `xml:"GetAllUserTransactions"`
	Token   string   `xml:"token,omitempty"`
}

type GetAllUserTransactionsResponse struct {
	XMLName                      xml.Name            `xml:"GetAllUserTransactionsResponse"`
	GetAllUserTransactionsResult *ArrayOfTransaction `xml:"GetAllUserTransactionsResult,omitempty"`
}

type ArrayOfTransaction struct {
	Transactions []*DTO.TransactionDTO `xml:"Transaction,omitempty"`
}

func (rsl RSLoyaltyWebService) GetUserAllTransactions(token string) ([]*DTO.TransactionDTO, error) {

	var request = &GetAllUserTransactions{Token: token}
	var response = &GetAllUserTransactionsResponse{}

	err := rsl.Soap("", request, "GetAllUserTransactions", response)
	if err != nil {
		return nil, err
	}

	return response.GetAllUserTransactionsResult.Transactions, nil
}

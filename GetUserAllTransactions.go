package methods

import (
	"encoding/xml"
	"log"
	"unibot/RSLoyatyWebService/DTO"
)

type GetAllUserTransactions struct {
	XMLName xml.Name `xml:" GetAllUserTransactions"`
	Token   string   `xml:"token,omitempty"`
}

type GetAllUserTransactionsResponse struct {
	XMLName                      xml.Name            `xml:"GetAllUserTransactionsResponse"`
	GetAllUserTransactionsResult *ArrayOfTransaction `xml:"GetAllUserTransactionsResult,omitempty"`
}

type ArrayOfTransaction struct {
	Transactions []*DTO.TransactionDTO `xml:"Transaction,omitempty"`
}

func (rsl RSLoyaltyWebService) GetUserAllTransactions(config RSLoyaltyWebService, channel string, identifier int64, token string) []*DTO.TransactionDTO {

	rsl.Init(config, channel, identifier)

	var request = &GetAllUserTransactions{Token: token}
	var response = &GetAllUserTransactionsResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "GetAllUserTransactions", response)
	if err != nil {
		log.Println("Ошибка выполнения запроса", err)
	}

	return response.GetAllUserTransactionsResult.Transactions
}

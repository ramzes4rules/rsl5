package methods

import (
	"encoding/xml"
	"log"
	"unibot/RSLoyatyWebService/DTO"
)

type GetMessages struct {
	XMLName   xml.Name `xml:"GetMessages"`
	Token     string   `xml:"token,omitempty"`
	CompanyId int64    `xml:"companyId,omitempty"`
}
type GetMessagesResponse struct {
	XMLName           xml.Name                  `xml:"GetMessagesResponse"`
	GetMessagesResult *ArrayOfMailingToCustomer `xml:"GetMessagesResult,omitempty"`
}
type ArrayOfMailingToCustomer struct {
	MailingToCustomer []*DTO.MailingToCustomerDTO `xml:"MailingToCustomer,omitempty"`
}

func (rsl RSLoyaltyWebService) GetMessages(config RSLoyaltyWebService, channel string, identifier int64, token string) []*DTO.MailingToCustomerDTO {

	rsl.Init(config, channel, identifier)

	var request = &GetMessages{Token: token}
	var response = &GetMessagesResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "GetMessages", response)
	if err != nil {
		log.Println("Ошибка выполнения запроса")
		//return settings
	}

	return response.GetMessagesResult.MailingToCustomer

}

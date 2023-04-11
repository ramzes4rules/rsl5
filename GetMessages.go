package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
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

func (rsl RSLoyaltyWebService) GetMessages(token string) ([]*DTO.MailingToCustomerDTO, error) {

	var request = &GetMessages{Token: token}
	var response = &GetMessagesResponse{}

	err := rsl.Soap("", request, "GetMessages", response)
	if err != nil {
		return nil, err
	}

	return response.GetMessagesResult.MailingToCustomer, nil

}

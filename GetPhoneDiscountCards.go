package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetPhoneDiscountCards struct {
	XMLName xml.Name `xml:"GetPhoneDiscountCards"`
	Token   string   `xml:"token,omitempty"`
	Phone   string   `xml:"phone,omitempty"`
}

type GetPhoneDiscountCardsResponse struct {
	XMLName                     xml.Name             `xml:"GetPhoneDiscountCardsResponse"`
	GetPhoneDiscountCardsResult *ArrayOfDiscountCard `xml:"GetPhoneDiscountCardsResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetPhoneDiscountCards(phone string) ([]*DTO.DiscountCardDTO, error) {

	var request = &GetPhoneDiscountCards{Phone: phone, Token: rsl.PrivateKey}
	var response = &GetPhoneDiscountCardsResponse{}

	err := rsl.Soap("", request, "GetPhoneDiscountCards", response)
	if err != nil {

		return nil, err
	}

	return response.GetPhoneDiscountCardsResult.DiscountCard, nil
}

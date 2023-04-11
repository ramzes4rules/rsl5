package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetDiscountCards struct {
	XMLName xml.Name `xml:"GetDiscountCards"`
	Token   string   `xml:"token,omitempty"`
}

type GetDiscountCardsResponse struct {
	XMLName                xml.Name             `xml:"GetDiscountCardsResponse"`
	GetDiscountCardsResult *ArrayOfDiscountCard `xml:"GetDiscountCardsResult,omitempty"`
}

type ArrayOfDiscountCard struct {
	DiscountCard []*DTO.DiscountCardDTO `xml:"DiscountCard,omitempty"`
}

func (rsl RSLoyaltyWebService) GetDiscountCards(token string) ([]*DTO.DiscountCardDTO, error) {

	var request = &GetDiscountCards{Token: token}
	var response = &GetDiscountCardsResponse{}

	err := rsl.Soap("", request, "GetDiscountCards", response)
	if err != nil {
		return nil, err
	}

	return response.GetDiscountCardsResult.DiscountCard, nil
}

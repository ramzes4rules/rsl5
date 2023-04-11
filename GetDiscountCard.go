package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetDiscountCard struct {
	XMLName            xml.Name `xml:"GetDiscountCard"`
	DiscountCardNumber string   `xml:"discountCardNumber,omitempty"`
}

type GetDiscountCardResponse struct {
	XMLName               xml.Name             `xml:"GetDiscountCardResponse"`
	GetDiscountCardResult *DTO.DiscountCardDTO `xml:"GetDiscountCardResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetDiscountCard(DiscountCardNumber string) (*DTO.DiscountCardDTO, error) {

	var request = &GetDiscountCard{DiscountCardNumber: DiscountCardNumber}
	var response = &GetDiscountCardResponse{}

	err := rsl.Soap("", request, "GetDiscountCard", response)
	if err != nil {
		return nil, err
	}

	return response.GetDiscountCardResult, nil

}

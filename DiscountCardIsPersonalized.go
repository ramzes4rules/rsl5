package methods

import (
	"encoding/xml"
)

type DiscountCardIsPersonalized struct {
	XMLName            xml.Name `xml:"DiscountCardIsPersonalized"`
	DiscountCardNumber string   `xml:"discountCardNumber,omitempty"`
}

type DiscountCardIsPersonalizedResponse struct {
	XMLName                          xml.Name `xml:"DiscountCardIsPersonalizedResponse"`
	DiscountCardIsPersonalizedResult bool     `xml:"DiscountCardIsPersonalizedResult,omitempty"`
}

func (rsl RSLoyaltyWebService) DiscountCardIsPersonalized(DiscountCardNumber string) (bool, error) {

	var request = &DiscountCardIsPersonalized{DiscountCardNumber: DiscountCardNumber}
	var response = &DiscountCardIsPersonalizedResponse{}

	err := rsl.Soap("", request, "DiscountCardIsPersonalized", response)
	if err != nil {
		return false, err
	}

	return response.DiscountCardIsPersonalizedResult, nil

}

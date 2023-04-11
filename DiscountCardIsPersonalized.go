package methods

import (
	"encoding/xml"
	"fmt"
)

type DiscountCardIsPersonalized struct {
	XMLName            xml.Name `xml:"DiscountCardIsPersonalized"`
	DiscountCardNumber string   `xml:"discountCardNumber,omitempty"`
}

type DiscountCardIsPersonalizedResponse struct {
	XMLName                          xml.Name `xml:"DiscountCardIsPersonalizedResponse"`
	DiscountCardIsPersonalizedResult bool     `xml:"DiscountCardIsPersonalizedResult,omitempty"`
}

func (rsl RSLoyaltyWebService) DiscountCardIsPersonalized(config RSLoyaltyWebService, channel string, identifier int64, DiscountCardNumber string) bool {

	rsl.Init(config, channel, identifier)

	var name = fmt.Sprintf("[%s/%d %s]", channel, identifier, "RSLoyaltyWebService.DiscountCardIsPersonalized")
	var request = &DiscountCardIsPersonalized{DiscountCardNumber: DiscountCardNumber}
	var response = &DiscountCardIsPersonalizedResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "DiscountCardIsPersonalized", response)
	if err != nil {
		nestor.Error(name, fmt.Sprintf("Ошибка выполнения запроса: %v", err))
	}

	return response.DiscountCardIsPersonalizedResult

}

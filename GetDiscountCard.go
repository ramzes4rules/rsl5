package methods

import (
	"encoding/xml"
	"fmt"
	"unibot/RSLoyatyWebService/DTO"
)

type GetDiscountCard struct {
	XMLName            xml.Name `xml:"GetDiscountCard"`
	DiscountCardNumber string   `xml:"discountCardNumber,omitempty"`
}

type GetDiscountCardResponse struct {
	XMLName               xml.Name             `xml:"GetDiscountCardResponse"`
	GetDiscountCardResult *DTO.DiscountCardDTO `xml:"GetDiscountCardResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetDiscountCard(channel string, identifier int64, DiscountCardNumber string) *DTO.DiscountCardDTO {

	var name = fmt.Sprintf("[%s/%d %s]", channel, identifier, "RSLoyaltyWebService.GetDiscountCard")
	var request = &GetDiscountCard{DiscountCardNumber: DiscountCardNumber}
	var response = &GetDiscountCardResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "GetDiscountCard", response)
	if err != nil {
		nestor.Error(name, fmt.Sprintf("Ошибка выполнения запроса: %v", err))
	}

	return response.GetDiscountCardResult

}

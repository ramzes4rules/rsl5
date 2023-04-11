package methods

import (
	"encoding/xml"
	"log"
)

type HasFreeVirtualDiscountCardsRequest struct {
	XMLName xml.Name `xml:"HasFreeVirtualDiscountCards"`
}

type HasFreeVirtualDiscountCardsResponse struct {
	XMLName xml.Name `xml:"HasFreeVirtualDiscountCardsResponse"`
	Result  bool     `xml:"HasFreeVirtualDiscountCardsResult"`
}

func (rsl RSLoyaltyWebService) HasFreeVirtualDiscountCards(channel string, identifier int64) bool {

	var request = &HasFreeVirtualDiscountCardsRequest{}
	var response = &HasFreeVirtualDiscountCardsResponse{}

	status, err := rsl.Soap2(channel, identifier, request, "HasFreeVirtualDiscountCards", response)
	if err != nil {
		log.Println("Ошибка выполнения запроса")
		return false
	}
	if status == "200 OK" {
		return response.Result
	} else {
		return false
	}
}

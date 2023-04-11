package methods

import (
	"encoding/xml"
	"fmt"
	"unibot/RSLoyatyWebService/DTO"
)

type GetDiscountCards struct {
	XMLName xml.Name `xml:" GetDiscountCards"`
	Token   string   `xml:"token,omitempty"`
}

type GetDiscountCardsResponse struct {
	XMLName                xml.Name             `xml:" GetDiscountCardsResponse"`
	GetDiscountCardsResult *ArrayOfDiscountCard `xml:"GetDiscountCardsResult,omitempty"`
}

type ArrayOfDiscountCard struct {
	DiscountCard []*DTO.DiscountCardDTO `xml:"DiscountCard,omitempty"`
}

func (rsl RSLoyaltyWebService) GetDiscountCards(config RSLoyaltyWebService, channel string, identifier int64, token string) []*DTO.DiscountCardDTO {

	rsl.Init(config, channel, identifier)

	var name = fmt.Sprintf(nestor.HeaderFormat, channel, identifier, "RSLoyaltyWebService.GetDiscountCards")
	nestor.Debug(name, "*** Получение токена авторизации.")

	var request = &GetDiscountCards{Token: token}
	var response = &GetDiscountCardsResponse{}

	status, err := rsl.Soap2(channel, identifier, request, "GetDiscountCards", response)
	nestor.Debug(name, fmt.Sprintf("Статус выполнения запроса: '%s'", status))

	if err != nil {
		nestor.Error(name, fmt.Sprintf("Ошибка выполнения запроса: '%v'", err))
		return nil
	}

	return response.GetDiscountCardsResult.DiscountCard
}

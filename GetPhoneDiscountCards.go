package methods

import (
	"encoding/xml"
	"unibot/RSLoyatyWebService/DTO"
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

func (rsl RSLoyaltyWebService) GetPhoneDiscountCards(phone string) []*DTO.DiscountCardDTO {

	//rsl.Init(config, channel, identifier)

	//var name = fmt.Sprintf(lhm, channel, identifier, "RSLoyaltyWebService.GetPhoneDiscountCards")
	//nestor.Debug(name, "Получение списка дисконтных карт по номеру телефона.")

	var request = &GetPhoneDiscountCards{Phone: phone, Token: rsl.PrivateKey}
	var response = &GetPhoneDiscountCardsResponse{}
	//nestor.Debug(name, "Объекты созданы.")

	_, err := rsl.Soap2("", 0, request, "GetPhoneDiscountCards", response)
	if err != nil {
		//nestor.Error(name, fmt.Sprintf("Ошибка выполнения запроса: %v", err))
		return []*DTO.DiscountCardDTO{}
	}

	return response.GetPhoneDiscountCardsResult.DiscountCard
}

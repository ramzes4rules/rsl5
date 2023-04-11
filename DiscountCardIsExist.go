package methods

import "encoding/xml"

type DiscountCardIsExist struct {
	XMLName            xml.Name `xml:"DiscountCardIsExist"`
	DiscountCardNumber string   `xml:"discountCardNumber,omitempty"`
}

type DiscountCardIsExistResponse struct {
	XMLName                   xml.Name `xml:"DiscountCardIsExistResponse"`
	DiscountCardIsExistResult bool     `xml:"DiscountCardIsExistResult,omitempty"`
}

func (rsl RSLoyaltyWebService) DiscountCardIsExist(config RSLoyaltyWebService, channel string, identifier int64, DiscountCardNumber string) bool {

	rsl.Init(config, channel, identifier)

	var request = &DiscountCardIsExist{DiscountCardNumber: DiscountCardNumber}
	var response = &DiscountCardIsExistResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "DiscountCardIsExist", response)
	if err != nil {
		//log.Println("Ошибка выполнения запроса")
		//return settings
	}

	return response.DiscountCardIsExistResult
}
